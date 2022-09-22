/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a python script",
	Long:  `Run a python script in a container`,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, sErr := cmd.Flags().GetString("sourceFile")
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if sErr != nil {
			return sErr
		}
		runErr := runAction(containerName, envName, source)
		return runErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		source, _ := cmd.Flags().GetString("sourceFile")

		msg := fmt.Sprintf("%q run in %q container with %q environment", source, containerName, envName)
		environ.ShowMessage(environ.INFO, msg)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAction(containerName string, envName string, source string) error {
	stOut, err := environ.RunScript(containerName, envName, source)
	if err != nil {
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, stOut)

		return err
	} else {
		return nil
	}
}
