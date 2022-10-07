/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update packages to conda environment",

	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		packageName, pErr := cmd.Flags().GetString("packageName")
		homePath, hErr := cmd.Flags().GetString("homePath")
		if hErr != nil {
			return hErr
		}
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if pErr != nil {
			return pErr
		}
		err := updateAction(containerName, envName, packageName, homePath)
		return err
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		packageName, _ := cmd.Flags().GetString("packageName")
		msg := fmt.Sprintf("%q updated in %q environment in %q container", packageName, envName, containerName)
		environ.ShowMessage(environ.INFO, msg)
	},
}

func init() {
	packageCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func updateAction(containerName string, envName string, packageName string, homePath string) error {
	sOut, err := environ.UpdatePackage(containerName, envName, packageName, homePath)
	if err != nil {
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, sOut)
	}
	return err
}
