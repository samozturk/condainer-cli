/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"git.tazi.ai/samet/rte-cli/environ"
	"git.tazi.ai/samet/rte-cli/utils"
	"github.com/spf13/cobra"
)

// addZipEnvCmd represents the AddZipEnv command
var addZipEnvCmd = &cobra.Command{
	Use:   "addZipEnv",
	Short: "Add conda environment from a zip file",
	Long:  `Add zipped conda environment file to the specified container`,
	RunE: func(cmd *cobra.Command, args []string) error {
		containerName, cErr := cmd.Flags().GetString("container")
		source, sErr := cmd.Flags().GetString("sourceFile")
		hostBindPath, hErr := cmd.Flags().GetString("hostBindPath")
		if cErr != nil {
			return cErr
		}
		if sErr != nil {
			return sErr
		}
		if hErr != nil {
			return hErr
		}
		err := addZipEnvAction(containerName, source, hostBindPath)
		return err
	},
}

func init() {
	environCmd.AddCommand(addZipEnvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addZipEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addZipEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addZipEnvCmd.Flags().StringP("sourceFile", "f", "", "path of compressed package directory")
	addZipEnvCmd.Flags().StringP("hostBindPath", "b", "/tmp/envs", "where you bind the host to container, default: /tmp/envs")
	addZipEnvCmd.MarkPersistentFlagRequired("sourceFile")
	addZipEnvCmd.MarkPersistentFlagRequired("hostBindPath")
}

func addZipEnvAction(containerName string, source string, hostBindPath string) error {
	err := environ.AddZipEnv(containerName, source, hostBindPath)
	if err != nil {
		utils.ShowMessage(utils.ERROR, err.Error())
		return err
	} else {
		return err
	}

}

// func addZipEnvAction(containerName string, source string) error {
// 	utils.ShowMessage(utils.INFO, source)
// 	// Get file extension
// 	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
// 	// Get home directory
// 	homedir, hErr := os.UserHomeDir()
// 	if hErr != nil {
// 		log.Fatal(hErr)
// 	}
// 	// Fix python3.7, it doesnt have to be like that always.
// 	dest := fmt.Sprintf("%v/tmp/envs/", homedir)
// 	if fileExt == "zip" {

// 		utils.ShowMessage(utils.WARNING, dest)
// 		utils.UnzipSource(source, dest)
// 		return nil
// 	} else if fileExt == "tar" {
// 		utils.Untar(source, dest)
// 	}
// 	return nil
// }
