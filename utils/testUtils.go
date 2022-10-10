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

//refactored
// func CreateEnv(containerName string, envName string) {
// 	cCommand := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/%v python=3.7.10 pip'"
// 	cmdStr := fmt.Sprintf(cCommand, containerName, envName)
// 	infoMessage := fmt.Sprintf("Creating environment: %v", cmdStr)
// 	ShowMessage(WARNING, infoMessage)
// 	_, _, err := RunCommand(cmdStr)
// 	if err != nil {
// 		ShowMessage(ERROR, "Environment create failed.")
// 	} else {
// 		ShowMessage(INFO, fmt.Sprintf("%v Environment Created", envName))
// 	}
// }

//refactored
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

// no need for refactoring
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// passed for refactoring
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

// Refactored
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

//TODO: define an init function to start a container.
