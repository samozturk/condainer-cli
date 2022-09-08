/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// AddZipEnvCmd represents the AddZipEnv command
var AddZipEnvCmd = &cobra.Command{
	Use:   "AddZipEnv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		source, sErr := cmd.Flags().GetString("sourceFile")

		if cErr != nil {
			return cErr
		}
		if sErr != nil {
			return sErr
		}
		err := addZipEnvAction(containerName, source)
		return err
	},
}

func init() {
	environCmd.AddCommand(AddZipEnvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddZipEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddZipEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addZipEnvAction(containerName string, source string) error {
	environ.ShowMessage(environ.INFO, source)
	// Get file extension
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	// Get home directory
	homedir, hErr := os.UserHomeDir()
	if hErr != nil {
		log.Fatal(hErr)
	}
	// Fix python3.7, it doesnt have to be like that always.
	dest := fmt.Sprintf("%v/tmp/envs/", homedir)
	if fileExt == "zip" {

		environ.ShowMessage(environ.WARNING, dest)
		environ.UnzipSource(source, dest)
		return nil
	} else if fileExt == "tar" {
		environ.Untar(source, dest)
	}
	return nil
}
