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

// installReqCmd represents the installReq command
var installReqCmd = &cobra.Command{
	Use:   "installReq",
	Short: "Use requirements.txt file to install packages",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, sErr := cmd.Flags().GetString("requirementsFile")
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
		if sErr != nil {
			return sErr
		}
		addErr := AddFromTextAction(containerName, envName, source, homePath)
		return addErr
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
	packageCmd.AddCommand(installReqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installReqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installReqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	installReqCmd.Flags().StringP("sourceFile", "f", "requirements.txt", "path of the file")
	installReqCmd.MarkFlagRequired("sourceFile")

}

func AddFromTextAction(containerName string, envName string, source string, homePath string) error {
	stdOut, stderr, err := pkg.AddFromText(containerName, envName, source, homePath)
	if err != nil {
		utils.ShowMessage(utils.ERROR, fmt.Sprintf("stdout: %v \n stderr: %v", stdOut, stderr))
		return err
	} else {
		return nil
	}
}
