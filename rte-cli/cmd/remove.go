/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove packages from a conda environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		packageName, pErr := cmd.Flags().GetString("packageName")
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if pErr != nil {
			return pErr
		}
		err := removeAction(containerName, envName, packageName)
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
	packageCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func removeAction(containerName string, envName string, packageName string) error {
	sOut, err := environ.RemovePackage(containerName, envName, packageName)
	if err != nil {
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, sOut)
	}
	return err
}
