package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const TestContainerUrl string = "registry.tazi.ai/base:2.0"

// UTILITY FUNCTIONS //
//refactored
func RunContainer(containerName string) {
	// Run a container to test commands
	rCommand := fmt.Sprintf("docker run -t --rm -d -v $HOME/tmp/envs:/home/tazi/miniconda3/envs --name %v %v /bin/bash", containerName, TestContainerUrl)
	ShowMessage(WARNING, rCommand)
	_, _, err := RunCommand(rCommand)
	// Error Handling //
	if err != nil {
		ShowMessage(ERROR, "Container run failed.")
	} else {
		ShowMessage(INFO, "Container Created")
	}
}

//refactored
func StopContainer(containerName string) {
	// Delete container if exists
	sCommand := fmt.Sprintf("docker stop %v", containerName)
	_, _, err := RunCommand(sCommand)
	if err != nil {
		ShowMessage(ERROR, "Container delete failed.")
	} else {
		ShowMessage(INFO, "Container Deleted")
	}
}

// Create environment for testing
func CreateTestEnv(containerName string, envName string, pythonVersion string, homePath string) (string, error) {
	maj, min, _ := GetPyVersion(pythonVersion)
	shortCmd := fmt.Sprintf("%v/miniconda3/bin/conda create -y -p %v/miniconda3/envs/%v python=%d.%d pip", homePath, homePath, envName, maj, min)
	cmdStr := fmt.Sprintf("docker exec %v /bin/bash -c %q", containerName, shortCmd)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, stderr, err := RunCommand(cmdStr)
	if err == nil {
		fmt.Println(out, stderr)
		ShowMessage(INFO, "Environment Created.")
	}
	return out, err
}

func CleanEnv(containerName string, envName string) {
	// Clean environment after testing for sanitation
	cCommand := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda env remove -n %v'"
	cmdStr := fmt.Sprintf(cCommand, containerName, envName)
	infoMessage := fmt.Sprintf("Removing environment: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	_, _, err := RunCommand(cmdStr)
	if err != nil {
		ShowMessage(ERROR, "Environment delete failed.")
	} else {
		ShowMessage(INFO, fmt.Sprintf("%v Environment Deleted", envName))
	}
}
func AddTestPackage(containerName string, envName string, packageName string, homePath string) (string, string, error) {
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, stderr, err := RunCommand(cmdStr)
	if err == nil && stderr == "" {
		fmt.Println(out)
		ShowMessage(INFO, fmt.Sprintf("Added package %q to %q environment in %q container.", packageName, envName, containerName))
		return out, stderr, err
	} else {
		ShowMessage(ERROR, stderr)
	}
	return out, stderr, err
}
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetExistingEnvNames(containerName string) []string {
	envs := make([]string, 1)
	command := fmt.Sprintf("docker exec %v bash -c '/home/tazi/miniconda3/bin/conda env list'", containerName)
	out, _ := exec.Command("/bin/sh", "-c", command).Output()
	sOut := fmt.Sprintf("%s", out)

	r, err := regexp.Compile(".+\\s{3,}")
	if err != nil {
		log.Fatal(err)
	}
	result := r.FindAllStringSubmatch(sOut, -1)

	for idx := range result {
		envs = append(envs, strings.TrimSpace(result[idx][0]))
	}
	return envs
}

func GetExistingPackageNames(containerName string, envName string) []string {
	envs := make([]string, 1)
	command := "docker exec %v bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip list'"
	cmdStr := fmt.Sprintf(command, containerName, envName)
	out, _, _ := RunCommand(cmdStr)
	sOut := fmt.Sprintf("%s", out)

	r, err := regexp.Compile(".+\\s{3,}")
	if err != nil {
		log.Fatal(err)
	}
	result := r.FindAllStringSubmatch(sOut, -1)

	for idx := range result {
		envs = append(envs, strings.TrimSpace(result[idx][0]))
	}
	idx := GetIndex(envs, "Package")
	return envs[idx+1:]
}

func GetIndex(slice []string, searchedValue string) int {
	for idx, value := range slice {
		if searchedValue == value {
			return idx
		}
	}
	return -1
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func DeletePackage(packageName string, envName string) {
	homedir, hErr := os.UserHomeDir()
	if hErr != nil {
		log.Fatal(hErr)
	}
	dest := fmt.Sprintf("%v/tmp/envs/%v/lib/python3.7/site-packages", homedir, envName)
	cmdStr := fmt.Sprintf("rm -r %v/%v", dest, packageName)
	cOut, cErr := exec.Command("/bin/sh", "-c", cmdStr).Output()
	cOutStr := fmt.Sprintf("%s", cOut)
	if cErr != nil {
		log.Println(cOutStr, '\n', cErr.Error())
	}
}
