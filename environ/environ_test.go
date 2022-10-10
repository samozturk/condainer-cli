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

func init() {
	utils.RunCommand("rm -rf $HOME/tmp/envs/*")
	utils.StopContainer(containerName)
	utils.RunContainer(containerName)
	utils.CleanEnv(containerName, envName)
	utils.ShowMessage(utils.INFO, "--Init function executed.--")

}

//** ENVIRONMENT TESTS **//
func TestCreateEnv(t *testing.T) {

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
		t.Errorf("%v is not in %v. Existing envs: %v", envName, containerName, envs)
	}

	/* Sanitation */
	utils.CleanEnv(containerName, envName)
}

func TestCloneEnv(t *testing.T) {

	// Create an environment to clone later
	utils.CreateTestEnv(containerName, envName, pythonVersion, homePath)

	/* Apply function */
	// Clone environment
	cOut, cErr := CloneEnv(containerName, envName, cloneEnvName, homePath)
	log.Println(cOut, '\n', cErr)

	/* Testing */
	// Get existing environment names and check <envName> is in them
	envs := utils.GetExistingEnvNames(containerName)
	if !(utils.StringInSlice(cloneEnvName, envs)) {
		t.Errorf("%v is not in %v. Existing envs: %v", cloneEnvName, containerName, envs)
	}

	/* Sanitation */
	utils.CleanEnv(containerName, envName)
	utils.CleanEnv(containerName, cloneEnvName)
}

func TestRemoveEnv(t *testing.T) {

	// Create an environment to remove later
	utils.CreateTestEnv(containerName, envName, pythonVersion, homePath)

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
		t.Errorf("%v couldn't be removed in %v. Existing envs: %v", envName, containerName, envs)
	}
	/* Sanitation */
	// utils.CleanEnv(containerName, envName)
}
