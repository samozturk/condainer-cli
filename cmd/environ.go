/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// environCmd represents the environ command
var environCmd = &cobra.Command{
	Use:     "environ",
	Aliases: []string{"e"},
	Short:   "Manage conda environments",
	Long: `Parent command to host subcommands for managing conda environments.
	Conda environments can be created, cloned and copied from a zip file.`,
}

func init() {
	rootCmd.AddCommand(environCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	environCmd.PersistentFlags().StringP("newEnvName", "n", "", "environment name for cloning a new environment")
	environCmd.PersistentFlags().StringP("pythonVersion", "v", "3.8", "Python version")
}
