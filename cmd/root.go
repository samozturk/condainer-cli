/*
Copyright © 2022 Samet Ozturk <samet@tazi.ai>

*/
/*
DONE: Create Env
DONE: Clone Env

DONE: Add Package
DONE: Update Package
DONE : Remove Package

TODO: Add Package From a File

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "rte-cli",
	Short:   "A brief description of your application",
	Version: "1.0.0",
	Long: `rte-cli is a utility for handling python runtime environment in containers for tazi. 
	
It allows you to create a conda environment, delete it, clone it and add/remove packages to/from specified environment.

You can customize environment or package using a command line flag.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rte-cli.yaml)")
	rootCmd.PersistentFlags().StringP("container", "c", "", "container name")
	rootCmd.PersistentFlags().StringP("envName", "e", "base", "conda environment name, default: base")
	rootCmd.PersistentFlags().StringP("packageName", "p", "", "package name")
	rootCmd.PersistentFlags().StringP("newEnvName", "n", "", "environment name for cloning a new environment")
	rootCmd.PersistentFlags().StringP("sourceFile", "f", "", "path of compressed package directory")
	rootCmd.PersistentFlags().StringP("requirementsFile", "r", "", "path of requirements txt")
	rootCmd.PersistentFlags().StringP("pythonVersion", "v", "3.8", "Python version")
	rootCmd.PersistentFlags().StringP("homePath", "h", "/home/tazi", "home path for container, default: /home/tazi")
	rootCmd.PersistentFlags().StringP("hostBindPath", "b", "/tmp/envs", "where you bind the host to container, default: /tmp/envs")

	rootCmd.MarkPersistentFlagRequired("container")
	rootCmd.MarkPersistentFlagRequired("envName")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
