/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a conda environment in a container",
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		pythonVersion, vErr := cmd.Flags().GetString("pythonVersion")
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if vErr != nil {
			return vErr
		}
		createErr := createAction(containerName, envName, pythonVersion, args)
		return createErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		pythonVersion, _ := cmd.Flags().GetString("pythonVersion")

		msg := fmt.Sprintf("%q environment created in %q container with python version %q", envName, containerName, pythonVersion)
		environ.ShowMessage(environ.INFO, msg)
	},
}

func init() {
	environCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createAction(containerName string, envName string, pythonVersion string, args []string) error {
	stOut, err := environ.CreateEnv(containerName, envName, pythonVersion)
	if err != nil {
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, stOut)

		return err
	} else {
		return nil
	}
}
