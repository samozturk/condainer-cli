/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// installReqCmd represents the installReq command
var installReqCmd = &cobra.Command{
	Use:   "installReq",
	Short: "Use requirements.txt file to install packages",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, sErr := cmd.Flags().GetString("requirementsFile")
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if sErr != nil {
			return sErr
		}
		cloneErr := AddFromTextAction(containerName, envName, source)
		return cloneErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		source, _ := cmd.Flags().GetString("requirementsFile")
		msg := fmt.Sprintf("packages have installed from %v to % venvironment in %v container", source, envName, containerName)
		environ.ShowMessage(environ.INFO, msg)
	},
}

func init() {
	packageCmd.AddCommand(installReqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installReqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installReqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func AddFromTextAction(containerName string, envName string, source string) error {
	stdOut, err := environ.AddFromText(containerName, envName, source)
	if err != nil {
		environ.ShowMessage(environ.ERROR, stdOut)
		return err
	} else {
		return nil
	}
}
