/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"git.tazi.ai/samet/rte-cli/environ"
	"github.com/spf13/cobra"
)

// addZipCmd represents the addZip command
var addZipCmd = &cobra.Command{
	Use:   "addZip",
	Short: "Add a package to a environment from compressed file",
	Long:  "### Python packages ### \n Python packages can be installed manually. For that, wheel files needed which can be found at https://pypi.org/simple/<package_name>  Simply find appropriate version and architecture.\n Or pip can do it for you. ```python -m pip download --only-binary :all: --dest . --no-cache <package_name> ``` \n Downloaded file can tarbal or wheel. If that's a tarball, untar it and find the setup.py \n in setup.py, you can dependency files with regex ```install_requires=\\[:*(.*?)\\]``` \n  If that's a whl; first unzip the whl file. Whl files are pretty much like zip files. It will yield two folders: one is named as <package_name> other is <{package_name}-{version}.dist-info>. Inside the latter, there is a file called METADATA. You can parse neccessary libraries from METADA using this regex ```Requires-Dist:\\s:*(.*?)\n```. All dependencies needs to be installed before installing the package. Also you need to install dependencies of dependencies and so on.",
	RunE: func(cmd *cobra.Command, args []string) error {
		//containerName, cErr := cmd.Flags().GetString("container")
		envName, eErr := cmd.Flags().GetString("envName")
		source, pErr := cmd.Flags().GetString("sourceFile")
		pythonVersion, vErr := cmd.Flags().GetString("pythonVersion")

		if vErr != nil {
			return vErr
		}
		if eErr != nil {
			return eErr
		}
		if pErr != nil {
			return pErr
		}
		err := AddZipAction(envName, source, pythonVersion)
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

func AddZipAction(envName string, source string, pythonVersion string) error {
	err := environ.AddZipPackage(envName, source, pythonVersion)
	if err != nil {
		environ.ShowMessage(environ.ERROR, err.Error())
	}
	return err
}
