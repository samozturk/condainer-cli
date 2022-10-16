package pkg

import (
	"fmt"
	"os"
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

// func AddZipPackage(envName string, source string, pythonVersion string) error {
// 	// Get file extension
// 	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
// 	// Get home directory
// 	homedir, hErr := os.UserHomeDir()
// 	if hErr != nil {
// 		log.Fatal(hErr)
// 	}

// 	major, minor, _ := utils.GetPyVersion(pythonVersion)

// 	dest := fmt.Sprintf("%v/tmp/envs/%v/lib/python%d.%d/site-packages", homedir, envName, major, minor)
// 	if fileExt == "zip" {
// 		utils.ShowMessage(utils.WARNING, dest)
// 		utils.UnzipSource(source, dest)
// 		return nil
// 	} else if fileExt == "tar" {
// 		utils.ShowMessage(utils.WARNING, dest)
// 		utils.Untar(source, dest)
// 	}
// 	return nil
// }

func AddZipPackages(containerName string, envName string, homePath string, source string) (string, string, error) {

	ext := filepath.Ext(source)
	sourceFileName := strings.TrimSuffix(filepath.Base(source), ext) // ??
	// Checkings
	envs := utils.GetExistingEnvNames(containerName)
	if !utils.StringInSlice(envName, envs) {
		utils.ShowMessage(utils.ERROR, fmt.Sprintf("%v not present in current environments:%v", envName, envs))
		os.Exit(1)
	}

	// Copy zipped wheelfiles to container
	copyCmd := fmt.Sprintf("docker cp %v %v:%v/%v%v", source, containerName, homePath, sourceFileName, ext)
	utils.ShowMessage(utils.WARNING, fmt.Sprintf("Running the command: %v", copyCmd))
	out, stderr, err := utils.RunCommand(copyCmd)
	if err != nil && stderr != "" {
		return out, stderr, err
	}
	installCmd := fmt.Sprintf("pip install --no-index --find-links %v/%v/ -r %v/%v/requirements.txt", homePath, sourceFileName, homePath, sourceFileName)
	command := fmt.Sprintf("docker exec %v bash -c 'unzip %v/%v; %v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; %v'",
		containerName, homePath, sourceFileName, homePath, homePath, envName, installCmd)
	utils.ShowMessage(utils.WARNING, fmt.Sprintf("Running the command: %v", command))
	out, stderr, err = utils.RunCommand(command)
	return out, stderr, err

}

func AddFromText(containerName string, envName string, source string, homePath string) (string, string, error) {
	fileName := filepath.Base(source)
	envs := utils.GetExistingEnvNames(containerName)
	if !utils.StringInSlice(envName, envs) {
		utils.ShowMessage(utils.ERROR, fmt.Sprintf("%v not present in current environments:%v", envName, envs))
		os.Exit(1)
	}
	cpCmd := fmt.Sprintf("docker cp %v %v:%v/%v", source, containerName, homePath, fileName)
	_, _, err := utils.RunCommand(cpCmd)
	if err != nil {
		utils.ShowMessage(utils.ERROR, "Couldn't copy requirements file to container.")
	}
	// activate environment name and execute pip install requirements.txt
	cmdStr := fmt.Sprintf("docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install -r %v/%v'",
		containerName, homePath, homePath, envName, homePath, fileName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)
	out, stderr, err := utils.RunCommand(cmdStr)

	return out, stderr, err
}

func GetPkgsFromContainer(containerName string, envName string, homePath string, dest string) (string, string, error) {
	// If you get conda version error, use this: conda update -n base -c defaults conda
	var endCmd string
	if envName == "base" {
		endCmd = fmt.Sprintf("%v/miniconda3/bin/pip list --format=freeze > requirements.txt; %v/miniconda3/bin/pip download -r requirements.txt -d wheelhouse; mv %v/requirements.txt %v/wheelhouse/requirements.txt; zip -r wheelhouse.zip wheelhouse",
			homePath, homePath, homePath, homePath)
	} else {
		endCmd = fmt.Sprintf("pip list --format=freeze > requirements.txt; pip download -r requirements.txt -d wheelhouse; mv %v/requirements.txt %v/wheelhouse/requirements.txt; zip -r wheelhouse.zip wheelhouse",
			homePath, homePath)
	}

	cmdStr := fmt.Sprintf(
		"docker exec %v bash -c 'source activate %v; %v'",
		containerName, envName, endCmd)
	out, stderr, err := utils.RunCommand(cmdStr)
	cmdStr = fmt.Sprintf("mkdir %v ; docker cp %v:%v/wheelhouse.zip %v/wheelhouse.zip", dest, containerName, homePath, dest)
	out, stderr, err = utils.RunCommand(cmdStr)
	return out, stderr, err
}

func GetPkgsFromHost(envName string, dest string) {
	_, _, err := utils.GetReqFileFromLocalEnv(envName, dest)
	if err != nil {
		os.Exit(1)
	}
	cmdStr := fmt.Sprintf("mkdir %v/wheelhouse ; pip download -r %v/requirements.txt -d %v/wheelhouse", dest, dest, dest)
	utils.RunCommand(cmdStr)
	utils.ZipWriter("%v/wheelhose", "%v/wheelhouse.zip")

}
