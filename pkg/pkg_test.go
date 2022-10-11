package pkg

import (
	"log"
	"testing"

	"git.tazi.ai/samet/rte-cli/utils"
)

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
	aOut, aErr := AddPackage(containerName, envName, packageName, homePath)
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

// func TestAddZipPackageAction(t *testing.T) {
// 	containerName := "tazitest"
// 	envName := "testenv"
// 	source := "sample_package.zip"
// 	packageName := FilenameWithoutExtension(source)
// 	utils.ShowMessage(utils.INFO, packageName)

// 	/* Prepare */

// 	/* Apply function */
// 	aErr := AddZipPackage(envName, source, pythonVersion)
// 	if aErr != nil {
// 		log.Println(aErr.Error())
// 	}
// 	/* Testing */
// 	// Get existing environment names and check <envName> is in them
// 	packages := GetExistingPackageNames(containerName, envName)
// 	if !(StringInSlice(packageName, packages)) {
// 		log.Fatalln(ErrEnvNotFound)
// 	}
// 	// defer this
// 	/* Sanitation */
// 	cleanEnv(containerName, envName)
// 	deleteContainer(containerName)
// }
