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

// addZipCmd represents the addZip command
var addZipCmd = &cobra.Command{
	Use:   "addZip",
	Short: "Add a package to a environment from compressed file",

	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, pErr := cmd.Flags().GetString("sourceFile")
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if pErr != nil {
			return pErr
		}
		err := addZipAction(containerName, envName, source)
		return err
	},
}

func init() {
	packageCmd.AddCommand(addZipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addZipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addZipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addZipAction(containerName string, envName string, source string) error {
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	homedir, hErr := os.UserHomeDir()
	if hErr != nil {
		log.Fatal(hErr)
	}
	dest := fmt.Sprintf("%v/tmp/envs/%v/lib/python3.7/site-packages", homedir, envName)
	if fileExt == "zip" {

		environ.ShowMessage(environ.WARNING, dest)
		environ.UnzipSource(source, dest)
		return nil
	} else if fileExt == "tar" {
		environ.Untar(source, dest)
	}
	return nil
}
