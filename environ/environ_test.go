package environ_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"testing"

	"git.tazi.ai/samet/rte-cli/environ"
)

var (
	ErrEnvNotFound = errors.New("specified env is not present in the container")
)

var (
	containerName = "tazitest"
	envName       = "testenv"
	cloneEnvName  = "clonetest"
	homePath      = "/home/tazi"
	pythonVersion = "3.8.3"
)

//** ENVIRONMENT TESTS **//
func TestCreateEnv(t *testing.T) {
	containerName := "tazitest"
	envName := "testenv"

	/* Prepare */
	deleteContainer(containerName)
	runContainer(containerName)

	/* Apply function */
	// Create environment
	_, crErr := environ.CreateEnv(containerName, envName, pythonVersion, homePath)
	// Error Handling //
	if crErr != nil {
		environ.ShowMessage(environ.ERROR, "Create environment failed.")
		environ.ShowMessage(environ.ERROR, crErr.Error())
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := GetExistingEnvNames(containerName)
	log.Println(envs)
	if !(StringInSlice(envName, envs)) {
		t.Error(ErrEnvNotFound)
	}

	/* Sanitation */
	cleanEnv(containerName, envName)
	deleteContainer(containerName)
}

func TestCloneEnv(t *testing.T) {
	// Variables

	/* Prepare */
	deleteContainer(containerName)
	runContainer(containerName)

	// Create an environment to clone later
	createEnv(containerName, envName)

	/* Apply function */
	// Clone environment
	cOut, cErr := environ.CloneEnv(containerName, envName, cloneEnvName, homePath)
	log.Println(cOut, '\n', cErr)

	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := GetExistingEnvNames(containerName)
	if !(StringInSlice(cloneEnvName, envs)) {
		log.Fatalln(ErrEnvNotFound)
	}

	/* Sanitation */
	cleanEnv(containerName, envName)
	cleanEnv(containerName, cloneEnvName)
	deleteContainer(containerName)

}

func TestRemoveEnv(t *testing.T) {
	containerName := "tazitest"
	envName := "testenv"

	/* Prepare */
	deleteContainer(containerName)
	runContainer(containerName)
	// Create an environment to remove later
	createEnv(containerName, envName)

	/* Apply function */
	_, crErr := environ.RemoveEnv(containerName, envName, homePath)
	// Error Handling //
	if crErr != nil {
		environ.ShowMessage(environ.ERROR, "Remove environment failed.")
		environ.ShowMessage(environ.ERROR, crErr.Error())
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := GetExistingEnvNames(containerName)
	if StringInSlice(envName, envs) {
		log.Fatalln(ErrEnvNotFound)
	}

	/* Sanitation */
	cleanEnv(containerName, envName)
	deleteContainer(containerName)
}

//** PACKAGE TESTS **//
func TestAddPackage(t *testing.T) {
	containerName := "tazitest"
	envName := "testenv"
	packageName := "numpy"

	/* Prepare */
	deleteContainer(containerName)
	runContainer(containerName)
	createEnv(containerName, envName)

	/* Apply function */
	aOut, aErr := environ.AddPackage(containerName, envName, packageName, homePath)
	if aErr != nil {
		log.Println(aOut)
	}

	/* Testing */
	// Get existing environment names and check <envName> is in them
	packages := GetExistingPackageNames(containerName, envName)
	if !(StringInSlice(packageName, packages)) {
		log.Fatalln(ErrEnvNotFound)
	}

	/* Sanitation */
	cleanEnv(containerName, envName)
	deleteContainer(containerName)
}

func TestAddZipPackageAction(t *testing.T) {
	containerName := "tazitest"
	envName := "testenv"
	source := "sample_package.zip"
	packageName := FilenameWithoutExtension(source)
	environ.ShowMessage(environ.INFO, packageName)

	/* Prepare */

	/* Apply function */
	aErr := environ.AddZipPackage(envName, source)
	if aErr != nil {
		log.Println(aErr.Error())
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	packages := GetExistingPackageNames(containerName, envName)
	if !(StringInSlice(packageName, packages)) {
		log.Fatalln(ErrEnvNotFound)
	}
	// defer this
	/* Sanitation */
	cleanEnv(containerName, envName)
	deleteContainer(containerName)
}

// TEST UTILITY FUNCTIONS
func TestGetPyVersion(t *testing.T) {
	exp_maj := 3
	exp_min := 8
	exp_patch := 3
	maj, min, patch := environ.GetPyVersion(pythonVersion)
	if maj != exp_maj {
		t.Errorf("Expected %d %T got %d %T", exp_maj, exp_maj, maj, maj)
	}
	if min != exp_min {
		t.Errorf("Expected %d %T got %d %T", exp_min, exp_min, min, min)
	}
	if patch != exp_patch {
		t.Errorf("Expected %d %T got %d %T", exp_patch, exp_patch, patch, patch)
	}

}

// UTILITY FUNCTIONS //
func deleteContainer(containerName string) {
	// Delete container if exists
	dCommand := fmt.Sprintf("docker rm -f %v", containerName)
	dOut, err := exec.Command("/bin/sh", "-c", dCommand).Output()
	dOutString := fmt.Sprintf("%s", dOut)
	if err != nil {
		// Print what went wrong
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, "Container delete failed.")
	}
	environ.ShowMessage(environ.INFO, "Container Deleted")
	log.Println(dOutString)
}

func runContainer(containerName string) {
	// Run a container to test commands
	rCommand := fmt.Sprintf("docker run -t --rm -d -v $HOME/tmp/envs:/home/tazi/miniconda3/envs --name %v registry.tazi.ai/tazi-rte:1.0.0 /bin/bash", containerName)
	_, err := exec.Command("/bin/sh", "-c", rCommand).Output()
	// Error Handling //
	if err != nil {
		// Print what went wrong
		environ.ShowMessage(environ.ERROR, err.Error())
		environ.ShowMessage(environ.ERROR, "Container run failed.")
	}
	environ.ShowMessage(environ.INFO, "Container Created")
}

func cleanEnv(containerName string, envName string) {
	// Clean environment after testing for sanitation
	cCommand := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda env remove -n %v'"
	cmdStr := fmt.Sprintf(cCommand, containerName, envName)
	infoMessage := fmt.Sprintf("Removing environment: %v", cmdStr)
	environ.ShowMessage(environ.WARNING, infoMessage)
	cOut, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	cOutstring := fmt.Sprintf("%s", cOut)
	log.Println(cOutstring)
}

func createEnv(containerName string, envName string) {
	command := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/%v python=3.7.10 pip'"
	cmdStr := fmt.Sprintf(command, containerName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	environ.ShowMessage(environ.WARNING, infoMessage)
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	log.Println(sOut)
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

	for idx, _ := range result {
		envs = append(envs, strings.TrimSpace(result[idx][0]))
	}
	return envs
}

func GetExistingPackageNames(containerName string, envName string) []string {
	envs := make([]string, 1)
	command := "docker exec taptazi bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate kernel-manager; pip list'"
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
	idx := getIndex(envs, "Package")
	return envs[idx+1:]
}

func getIndex(slice []string, searchedValue string) int {
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
