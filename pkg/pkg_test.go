package pkg

import (
	"fmt"
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
	packageName   = "numpy"
	source        = testPath + "/goldenFiles/wheelhouse.zip"
	zippedPackage = "black"
)

func init() {
	utils.RunCommand(fmt.Sprintf("rm -rf %v/*", hostBindPath))
	utils.StopContainer(containerName)
	time.Sleep(30 * time.Second)
	utils.RunContainer(containerName)
	utils.CreateTestEnv(containerName, envName, pythonVersion, homePath)
	utils.ShowMessage(utils.INFO, "--Init function executed.--")
}

// //** PACKAGE TESTS **//
// func TestAddPackage(t *testing.T) {
// 	/* Apply function */
// 	aOut, aErr := AddPackage(containerName, envName, packageName, homePath)
// 	if aErr != nil {
// 		t.Error(aOut, "--", aErr)
// 	}
// 	/* Testing */
// 	// Get existing environment names and check <envName> is in them
// 	packages := utils.GetExistingPackageNames(containerName, envName)
// 	if !(utils.StringInSlice(packageName, packages)) {
// 		t.Errorf("%q not present in %v", packageName, packages)
// 	}
// }

// func TestRemovePackage(t *testing.T) {
// 	utils.AddTestPackage(containerName, envName, packageName, homePath)
// 	rOut, rErr := RemovePackage(containerName, envName, packageName, homePath)
// 	if rErr != nil {
// 		t.Errorf(rOut, "--", rErr)
// 	}
// 	/* Testing */
// 	// Get existing environment names and check <envName> is in them
// 	packages := utils.GetExistingPackageNames(containerName, envName)
// 	if utils.StringInSlice(packageName, packages) {
// 		t.Errorf("%q didn't removed from in %v", packageName, packages)
// 	}

// }

func TestAddZipPackages(t *testing.T) {
	/* Apply function */
	out, stderr, err := AddZipPackages(containerName, envName, homePath, source)
	if err != nil {
		t.Error(out, "--", stderr)
	}
	/* Testing */
	// Get existing environment names and check <envName> is in them
	packages := utils.GetExistingPackageNames(containerName, envName)
	if !utils.StringInSlice(zippedPackage, packages) {
		t.Errorf("%q is not present in %v", zippedPackage, packages)
	}

}
