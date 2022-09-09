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
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Bomcat struct {
	impl bomcatImplementation
}

func New() *Bomcat {
	return &Bomcat{
		impl: &defaultBomcatImplementation{},
	}
}

func (bc *Bomcat) Install(repoSlug, _ string) error {
	ctx := context.Background()
	release, err := bc.impl.GetLatestRepositoryRelease(ctx, repoSlug)
	if err != nil {
		return fmt.Errorf("finding repository release: %w", err)
	}

	// Search for an SBOM int the release
	url, err := bc.impl.GetReleaseSBOM(ctx, repoSlug, release)
	if err != nil {
		return fmt.Errorf("getting sbom URL: %w", err)
	}

	// If not found, we can get exit here
	if url == "" {
		return errors.New("could not find an SBOM in the release")
	}

	logrus.Infof("downloading sbom from %s", url)

	// Look for binaries
	// Download binary
	// Install
	return nil
}
