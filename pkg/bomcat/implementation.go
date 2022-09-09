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

package bomcat

import (
	"context"

	"github.com/puerco/bomcat/pkg/github"
)

type bomcatImplementation interface {
	GetLatestRepositoryRelease(context.Context, string) (int64, error)
	GetReleaseSBOM(context.Context, string, int64) (string, error)
	FindBinary() error
	DownloadBinary() error
	InstallBinary() error
}

type defaultBomcatImplementation struct{}

func (di *defaultBomcatImplementation) GetLatestRepositoryRelease(ctx context.Context, repoSlug string) (id int64, err error) {
	github := github.New()
	return github.GetLatestRepositoryRelease(ctx, repoSlug)
}

// GetReleaseSBOM reads the assets of a release ID and if
// it finds an SBOM returns it parsed
func (di *defaultBomcatImplementation) GetReleaseSBOM(ctx context.Context, repoSlug string, releaseID int64) (string, error) {
	github := github.New()
	return github.GetReleaseSBOM(ctx, repoSlug, releaseID)
}
func (di *defaultBomcatImplementation) FindBinary() error     { return nil }
func (di *defaultBomcatImplementation) DownloadBinary() error { return nil }
func (di *defaultBomcatImplementation) InstallBinary() error  { return nil }
