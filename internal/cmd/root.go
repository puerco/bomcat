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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/log"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Short: "A tool for working with SPDX manifests",
		Long: `tejolote (the handle of a molcajete, where you make salsa)
	
🌶 tejolote is a utility that allows a developer to execute a 
process - ideally a builder - and record its inputs and outputs.
The main goal is to obtain provenance information of builds
and other transformations when building and shipping software.

In its simplest form, you can precede your existing build 
command with tejolote run and it will make its best to create a
meaningful attestation. For example:

	If your build command is:
	make build

	CHange it with:
	tejolote run make build

Tejolote will try to make sane asumptions but for best results, it
allows for full control of the process you run.
	
	`,
		Use:               "tejolote",
		SilenceUsage:      false,
		PersistentPreRunE: initLogging,
	}

	rootCmd.PersistentFlags().StringVar(
		&commandLineOpts.logLevel,
		"log-level",
		"info",
		fmt.Sprintf("the logging verbosity, either %s", log.LevelNames()),
	)

	addInstall(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}

type commandLineOptions struct {
	logLevel string
}

var commandLineOpts = &commandLineOptions{}

func initLogging(*cobra.Command, []string) error {
	return log.SetupGlobalLogger(commandLineOpts.logLevel)
}
