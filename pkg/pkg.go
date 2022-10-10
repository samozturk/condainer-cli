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
	var command string

	command = "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func RemovePackage(containerName string, envName string, packageName string, homePath string) (string, error) {
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip uninstall -y %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func UpdatePackage(containerName string, envName string, packageName string, homePath string) (string, error) {

	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v --upgrade'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
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
