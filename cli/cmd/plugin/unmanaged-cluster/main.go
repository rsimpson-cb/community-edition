// Copyright 2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/unmanaged-cluster/cmd"
	cliv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/cli/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/buildinfo"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/command/plugin"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/log"
)

var description = `Deploy and manage single-node, static, Tanzu clusters.`

var descriptor = cliv1alpha1.PluginDescriptor{
	Name:        "unmanaged-cluster",
	Aliases:     []string{"um", "uc", "unmanaged"},
	Description: description,
	Group:       cliv1alpha1.RunCmdGroup,
	Version:     buildinfo.Version,
}

var (
	// logLevel is the verbosity to print
	logLevel int32

	// Log file to dump logs to
	logFile string

	// default build version
	defaultVersion = "v0.0.1-unversioned"
)

func main() {
	if descriptor.Version == "" {
		descriptor.Version = defaultVersion
	}

	// plugin!
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "unable to initialize new plugin")
	}

	p.Cmd.PersistentFlags().Int32VarP(&logLevel, "verbose", "v", 0, "Number for the log level verbosity(0-9)")
	p.Cmd.PersistentFlags().StringVar(&logFile, "log-file", "", "Log file path")

	p.AddCommands(
		cmd.ConfigureCmd,
		cmd.CreateCmd,
		cmd.DeleteCmd,
		cmd.ListCmd,
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
