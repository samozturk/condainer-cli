/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/pkg"
	"git.tazi.ai/samet/rte-cli/utils"
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
		err := removeAction(containerName, envName, packageName, homePath)
		return err
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		packageName, _ := cmd.Flags().GetString("packageName")
		msg := fmt.Sprintf("%q removed from %q environment in %q container", packageName, envName, containerName)
		utils.ShowMessage(utils.INFO, msg)
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
	// removeCmd.Flags().StringP("packageName", "p", "", "package name")
	// removeCmd.MarkFlagRequired("packageName")
}
func removeAction(containerName string, envName string, packageName string, homePath string) error {
	sOut, err := pkg.RemovePackage(containerName, envName, packageName, homePath)
	if err != nil {
		utils.ShowMessage(utils.ERROR, err.Error())
		utils.ShowMessage(utils.ERROR, sOut)
	}
	return err
}
