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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Install packages to conda environment",
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
		err := addAction(containerName, envName, packageName, homePath)
		return err
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		packageName, _ := cmd.Flags().GetString("packageName")
		msg := fmt.Sprintf("%q installed in %q environment in %q container", packageName, envName, containerName)
		utils.ShowMessage(utils.INFO, msg)
	},
}

func init() {
	packageCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("packageName", "p", "", "package name")
	addCmd.MarkFlagRequired("packageName")
}

func addAction(containerName string, envName string, packageName string, homePath string) error {
	sOut, err := pkg.AddPackage(containerName, envName, packageName, homePath)
	if err != nil {
		utils.ShowMessage(utils.ERROR, err.Error())
		utils.ShowMessage(utils.ERROR, sOut)

	}
	return err
}
