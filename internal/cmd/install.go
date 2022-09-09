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

package cmd

import (
	"fmt"

	"github.com/puerco/bomcat/pkg/bomcat"
	"github.com/spf13/cobra"
)

type installOptions struct{}

func addInstall(parentCommand *cobra.Command) {
	installOpts := installOptions{}
	installCmd := &cobra.Command{
		Short: "Install a binary",
		Long: `bomcat install [repository-url]

The bomcat install subcommand reads a repository's releases and 
will fetch an sbom from its release assets. It will then look in the
sbom for a binary for the running platform and install it.
        
        `,
		Use:               "install",
		SilenceUsage:      false,
		PersistentPreRunE: initLogging,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := runInstall(args, installOpts); err != nil {
				return fmt.Errorf("installing from repo: %w", err)
			}
			return nil
		},
	}

	parentCommand.AddCommand(installCmd)

}

func runInstall(args []string, opts installOptions) error {
	bc := bomcat.New()
	binary := ""
	if len(args) > 1 {
		binary = args[1]
	}
	return bc.Install(args[0], binary)
}
