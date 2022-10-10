package environ

import (
	"errors"
	"log"
	"testing"

	"git.tazi.ai/samet/rte-cli/utils"
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
	utils.StopContainer(containerName)
	utils.RunContainer(containerName)

	/* Apply function */
	// Create environment
	_, crErr := CreateEnv(containerName, envName, pythonVersion, homePath)
	// Error Handling //
	if crErr != nil {
		utils.ShowMessage(utils.ERROR, "Create environment failed.")
		utils.ShowMessage(utils.ERROR, crErr.Error())
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := utils.GetExistingEnvNames(containerName)
	log.Println(envs)
	if !(utils.StringInSlice(envName, envs)) {
		t.Error(ErrEnvNotFound)
	}

	/* Sanitation */
	utils.CleanEnv(containerName, envName)
	utils.StopContainer(containerName)
}

func TestCloneEnv(t *testing.T) {
	// Variables

	/* Prepare */
	utils.StopContainer(containerName)
	utils.RunContainer(containerName)

	// Create an environment to clone later
	utils.CreateEnv(containerName, envName)

	/* Apply function */
	// Clone environment
	cOut, cErr := CloneEnv(containerName, envName, cloneEnvName, homePath)
	log.Println(cOut, '\n', cErr)

	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := utils.GetExistingEnvNames(containerName)
	if !(utils.StringInSlice(cloneEnvName, envs)) {
		log.Fatalln(ErrEnvNotFound)
	}

	/* Sanitation */
	utils.CleanEnv(containerName, envName)
	utils.CleanEnv(containerName, cloneEnvName)
	utils.StopContainer(containerName)

}

func TestRemoveEnv(t *testing.T) {
	containerName := "tazitest"
	envName := "testenv"

	/* Prepare */
	utils.StopContainer(containerName)
	utils.RunContainer(containerName)
	// Create an environment to remove later
	utils.CreateEnv(containerName, envName)

	/* Apply function */
	_, crErr := RemoveEnv(containerName, envName, homePath)
	// Error Handling //
	if crErr != nil {
		utils.ShowMessage(utils.ERROR, "Remove environment failed.")
		utils.ShowMessage(utils.ERROR, crErr.Error())
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := utils.GetExistingEnvNames(containerName)
	if utils.StringInSlice(envName, envs) {
		log.Fatalln(ErrEnvNotFound)
	}

	/* Sanitation */
	utils.CleanEnv(containerName, envName)
	utils.StopContainer(containerName)
}

// TEST UTILITY FUNCTIONS
func TestGetPyVersion(t *testing.T) {
	exp_maj := 3
	exp_min := 8
	exp_patch := 3
	maj, min, patch := utils.GetPyVersion(pythonVersion)
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
