package utils

import "testing"

var (
	containerName = "tazitest"
	envName       = "testenv"
	cloneEnvName  = "clonetest"
	homePath      = "/home/tazi"
	pythonVersion = "3.8.3"
)

// TEST UTILITY FUNCTIONS
func TestGetPyVersion(t *testing.T) {
	exp_maj := 3
	exp_min := 8
	exp_patch := 3
	maj, min, patch := GetPyVersion(pythonVersion)
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
