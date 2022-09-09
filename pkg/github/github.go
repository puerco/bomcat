/*
Copyright 2022 Adolfo García Veytia

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	gogithub "github.com/google/go-github/v47/github"
	"github.com/sirupsen/logrus"
)

type GitHub struct {
	client *gogithub.Client
}

func New() *GitHub {
	return &GitHub{
		client: gogithub.NewClient(nil),
	}

}

func parseSlug(repoSlug string) (org, repo string, err error) {
	org, repo, ok := strings.Cut(repoSlug, "/")
	if !ok {
		err = errors.New("repository slug not wel formed, must be org/repo")
	}
	return org, repo, err
}

// GetLatestRepositoryRelease finds the last release in a repository
func (github *GitHub) GetLatestRepositoryRelease(ctx context.Context, repoSlug string) (releaseID int64, err error) {
	org, repo, err := parseSlug(repoSlug)
	if err != nil {
		return releaseID, err
	}

	release, _, err := github.client.Repositories.GetLatestRelease(context.Background(), org, repo)
	releaseID = release.GetID()

	logrus.Debugf("Latest repository release ID is: %d", releaseID)

	// return release, resp, err
	return releaseID, err
}

// GetReleaseSBOM returns the URL to download the SBOM of a GitHub release
// or an empty string if it cannot be found
func (github *GitHub) GetReleaseSBOM(ctx context.Context, repoSlug string, releaseID int64) (url string, err error) {
	org, repo, err := parseSlug(repoSlug)
	if err != nil {
		return "", err
	}
	assets, _, err := github.client.Repositories.ListReleaseAssets(
		ctx, org, repo, releaseID, &gogithub.ListOptions{},
	)
	if err != nil {
		return "", fmt.Errorf("getting release assets from GitHub: %w", err)
	}

	for _, a := range assets {
		name := a.GetName()
		if strings.Contains(name, ".spdx") {
			return a.GetBrowserDownloadURL(), nil
		}
	}
	return "", nil
}
