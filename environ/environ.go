package environ

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.tazi.ai/samet/rte-cli/utils"
)

// Create environment in a container with specific version of python
func CreateEnv(containerName string, envName string, pythonVersion string, homePath string) (string, error) {
	maj, min, _ := utils.GetPyVersion(pythonVersion)
	shortCmd := fmt.Sprintf("%v/miniconda3/bin/conda create -y -p %v/miniconda3/envs/%v python=%d.%d pip", homePath, homePath, envName, maj, min)
	cmdStr := fmt.Sprintf("docker exec %v /bin/bash -c %q", containerName, shortCmd)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil {
		fmt.Println(out, stderr)
		utils.ShowMessage(utils.INFO, "Environment Created.")
	}
	return out, err
}

// Clone an existing environment in a container
func CloneEnv(containerName string, envName string, newEnvName string, homePath string) (string, error) {

	command := "docker exec %v /bin/bash -c '%v/miniconda3/bin/conda create -y --name %v --clone %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, newEnvName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil {
		fmt.Println(out, stderr)
		utils.ShowMessage(utils.INFO, "Environment Cloned.")
	}
	return out, err
}

// Remove an existing environment
func RemoveEnv(containerName string, envName string, homePath string) (string, error) {
	command := "docker exec %v /bin/bash -c '%v/miniconda3/bin/conda env remove -n %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	envs := utils.GetExistingEnvNames(containerName)
	utils.ShowMessage(utils.WARNING, infoMessage)
	if !utils.StringInSlice(envName, envs) {
		utils.ShowMessage(utils.ERROR, "Environment doesn't exist.")
		os.Exit(1)
	}
	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil {
		fmt.Println(out, stderr)
		utils.ShowMessage(utils.INFO, "Environment Removed.")
	}
	return out, err
}

// Add an environment from a comporessed file
func AddZipEnv(containerName string, source string, hostBindPath string) error {
	// Get file extension
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	dest := hostBindPath

	if fileExt == "zip" {
		utils.ShowMessage(utils.WARNING, fmt.Sprintf("Copying to %q", dest))
		err := utils.UnzipSource(source, dest)
		if err != nil {
			utils.ShowMessage(utils.ERROR, fmt.Sprintf("%q", err.Error()))
			utils.ShowMessage(utils.WARNING, "This subcommand is deprecated. Please use 'addZipPackages' subcommmand under 'package' command.")
		}
		return err
	} else if fileExt == "tar" {
		utils.ShowMessage(utils.WARNING, fmt.Sprintf("Copying to %q", dest))
		utils.Untar(source, dest)
		utils.ShowMessage(utils.WARNING, "This subcommand is deprecated. Please use 'addZipPackages' subcommmand under 'package' command.")
	} else {
		utils.ShowMessage(utils.ERROR, fmt.Sprintf("%v extension is not supported.", fileExt))
	}
	return nil
}
