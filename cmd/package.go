/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Manage python packages",
	Long: `Parent command to host subcommands for managing python packages.
You can install, update or remove python libraries.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("package called")
	// },
}

func init() {
	rootCmd.AddCommand(packageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	packageCmd.PersistentFlags().StringP("packageName", "p", "", "package name")
	packageCmd.MarkPersistentFlagRequired("packageName")

}
