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

// getZipCmd represents the getZip command
var getZipCmd = &cobra.Command{
	Use:   "getZip",
	Short: "Zip python packages of an environment for offline use",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, sErr := cmd.Flags().GetString("requirementsFile")
		homePath, hErr := cmd.Flags().GetString("homePath")
		local, lErr := cmd.Flags().GetBool("local")
		dest, dErr := cmd.Flags().GetString("destination")
		if hErr != nil {
			return hErr
		}
		if cErr != nil {
			return cErr
		}
		if eErr != nil {
			return eErr
		}
		if sErr != nil {
			return sErr
		}
		if lErr != nil {
			return lErr
		}
		if lErr != nil {
			return lErr
		}
		if dErr != nil {
			return dErr
		}
		cloneErr := getZipAction(containerName, envName, source, homePath, local, dest)
		return cloneErr
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("container")
		envName, _ := cmd.Flags().GetString("envName")
		source, _ := cmd.Flags().GetString("requirementsFile")
		msg := fmt.Sprintf("packages have installed from %v to %v environment in %v container", source, envName, containerName)
		utils.ShowMessage(utils.INFO, msg)
	},
}

func init() {
	packageCmd.AddCommand(getZipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getZipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getZipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getZipAction(containerName string, envName string, source string, homePath string, local bool, dest string) error {
	if local {
		stdout, stderr, err := pkg.GetPkgsFromHost(envName, dest)
		if err != nil {
			utils.ShowMessage(utils.ERROR, fmt.Sprintf("stdout: %v \n stderr: %v", stdout, stderr))
			return err
		} else {
			return nil
		}
	} else {
		stdout, stderr, err := pkg.GetPkgsFromContainer(containerName, envName, homePath, dest)
		if err != nil {
			utils.ShowMessage(utils.ERROR, fmt.Sprintf("stdout: %v \n stderr: %v", stdout, stderr))
			return err
		} else {
			return nil
		}
	}

}
