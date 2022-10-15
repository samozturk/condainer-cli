package runScript

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"git.tazi.ai/samet/rte-cli/utils"
)

var (
	containerName = "tazitest"
	envName       = "testenv"
	homePath      = "/home/tazi"
	pythonVersion = "3.8.3"
	testPath      = "/Users/samet/Documents/Projects/rte-cli"
	hostBindPath  = "/Users/samet/tmp/envs"
	scriptFile    = "goldenFiles/test_script.py"
)

func init() {
	utils.RunCommand(fmt.Sprintf("rm -rf %v/*", hostBindPath))
	utils.StopContainer(containerName)
	time.Sleep(15 * time.Second)
	utils.RunContainer(containerName)
	utils.CreateTestEnv(containerName, envName, pythonVersion, homePath)
	utils.ShowMessage(utils.INFO, "--Init function executed.--")
}

func TestRunPy(t *testing.T) {
	fileName := filepath.Base(scriptFile)
	cpCmd := fmt.Sprintf("docker cp %v/%v tazitest:%v/%v", testPath, scriptFile, homePath, fileName)
	_, _, err := utils.RunCommand(cpCmd)
	if err != nil {
		t.Errorf("Couldn't copy python file to container")
	}
	RunPy(containerName, envName, (homePath + "/" + fileName), homePath)

}
