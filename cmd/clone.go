/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"git.tazi.ai/samet/rte-cli/utils"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones a conda environment",

	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		newEnvName, nErr := cmd.Flags().GetString("newEnvName")
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
		if nErr != nil {
			return nErr
		}
		cloneErr := cloneAction(containerName, envName, newEnvName, homePath)
		return cloneErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		newEnvName, _ := cmd.Flags().GetString("newEnvName")
		msg := fmt.Sprintf("%v environment cloned as %v in %v container", envName, newEnvName, containerName)
		utils.ShowMessage(utils.INFO, msg)
	},
}

func init() {
	environCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cloneCmd.Flags().StringP("newEnvName", "n", "", "environment name for cloning a new environment")
	cloneCmd.MarkPersistentFlagRequired("newEnvName")
}

func cloneAction(containerName string, envName string, newEnvName string, homePath string) error {
	stdOut, err := environ.CloneEnv(containerName, envName, newEnvName, homePath)
	if err != nil {
		utils.ShowMessage(utils.ERROR, stdOut)
		return err
	} else {
		return nil
	}
}
