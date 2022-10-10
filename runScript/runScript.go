package runScript

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"git.tazi.ai/samet/rte-cli/utils"
)

func RunPy(containerName string, envName string, source string, homePath string) (string, error) {
	// Get file extension
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	// Check file format
	if fileExt != "py" {
		log.Fatalf("%q is not a python file.", source)
	}
	command := "docker exec %v bash -c '%v/miniconda3/bin/conda init; source %v/miniconda3/etc/profile.d/conda.sh; conda activate %v; python3 %v'"
	cmdStr := fmt.Sprintf(command, containerName, homePath, homePath, envName, source)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	utils.ShowMessage(utils.WARNING, infoMessage)

	out, stderr, err := utils.RunCommand(cmdStr)
	if err == nil {
		fmt.Println(out, stderr)
		utils.ShowMessage(utils.INFO, "Script Ran.")
	}
	return out, err

}
