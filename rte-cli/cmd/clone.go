/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
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
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if nErr != nil {
			return eErr
		}
		cloneErr := cloneAction(containerName, envName, newEnvName, args)
		return cloneErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		newEnvName, _ := cmd.Flags().GetString("newEnvName")
		msg := fmt.Sprintf("%v environment cloned as %v in %v container", envName, newEnvName, containerName)
		environ.ShowMessage(environ.INFO, msg)
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
}

func cloneAction(containerName string, envName string, newEnvName string, args []string) error {
	stdOut, err := environ.CloneEnv(containerName, envName, newEnvName)
	if err != nil {
		environ.ShowMessage(environ.ERROR, stdOut)
		return err
	} else {
		return nil
	}
}
