package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"git.tazi.ai/samet/rte-cli/utils"
)

const envBindDir string = "$HOME/tmp/envs3"

func AddPackage(containerName string, envName string, packageName string, homePath string) (string, error) {
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil && stderr == "" {
		fmt.Println(out)
		utils.ShowMessage(utils.INFO, fmt.Sprintf("Added package %q to %q environment in %q container.", packageName, envName, containerName))
		return out, err
	} else {
		utils.ShowMessage(utils.ERROR, stderr)
	}
	return out, err
}

func RemovePackage(containerName string, envName string, packageName string, homePath string) (string, error) {
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip uninstall -y %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil && stderr == "" {
		fmt.Println(out)
		utils.ShowMessage(utils.INFO, fmt.Sprintf("Removed package %q to %q environment in %q container.", packageName, envName, containerName))
		return out, err
	} else {
		utils.ShowMessage(utils.ERROR, stderr)
	}
	return out, err
}

func UpdatePackage(containerName string, envName string, packageName string, homePath string) (string, error) {

	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v --upgrade'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	out, stderr, err := utils.RunCommand(cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	if err == nil && stderr == "" {
		fmt.Println(out)
		utils.ShowMessage(utils.INFO, fmt.Sprintf("Updated package %q to %q environment in %q container.", packageName, envName, containerName))
		return out, err
	} else {
		utils.ShowMessage(utils.ERROR, stderr)
	}
	return out, err
}

func AddZipPackage(envName string, source string, pythonVersion string) error {
	// Get file extension
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	// Get home directory
	homedir, hErr := os.UserHomeDir()
	if hErr != nil {
		log.Fatal(hErr)
	}

	major, minor, _ := utils.GetPyVersion(pythonVersion)

	dest := fmt.Sprintf("%v/tmp/envs/%v/lib/python%d.%d/site-packages", homedir, envName, major, minor)
	if fileExt == "zip" {
		utils.ShowMessage(utils.WARNING, dest)
		utils.UnzipSource(source, dest)
		return nil
	} else if fileExt == "tar" {
		utils.ShowMessage(utils.WARNING, dest)
		utils.Untar(source, dest)
	}
	return nil
}

func AddFromText(containerName string, envName string, source string, homePath string, hostBindPath string) (string, error) {
	// Copy requirements.txt to miniconda3/envs/ which is a shared directory
	dst := fmt.Sprintf("%v/envs/%v/requirements.txt", hostBindPath, envName)
	utils.CopyFile(source, dst)
	// activate environment name and execute pip install requirements.txt √
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install -r %v/miniconda3/envs/%v/requirements.txt'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, homePath, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err

}

func GetPkgsFromContainer(containerName string, envName string, homePath string, dest string) (string, string, error) {
	// If you get conda version error, use this: conda update -n base -c defaults conda
	var endCmd string
	if envName == "base" {
		endCmd = fmt.Sprintf("%v/miniconda3/bin/pip list --format=freeze > requirements.txt; %v/miniconda3/bin/pip download -r requirements.txt -d wheelhouse; zip -r wheelhouse.zip wheelhouse",
			homePath, homePath)
	} else {
		endCmd = "pip list --format=freeze > requirements.txt; pip download -r requirements.txt -d wheelhouse; zip -r wheelhouse.zip wheelhouse"
	}

	cmdStr := fmt.Sprintf(
		"docker exec %v bash -c 'source activate %v; %v'",
		containerName, envName, endCmd)
	out, stderr, err := utils.RunCommand(cmdStr)
	cmdStr = fmt.Sprintf("docker cp %v:%v/wheelhouse.zip %v/wheelhouse.zip", containerName, homePath, dest)
	out, stderr, err = utils.RunCommand(cmdStr)
	return out, stderr, err
}

func GetPkgsFromHost() {

}
