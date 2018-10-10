package config

import (
	"testing"

	"github.com/joshuakwan/hydra/utils"
)

func TestDevStorageType(t *testing.T) {
	utils.AssertEqual(t, GetStorageType(), "etcdv3", "GetStorageType() failed")
}

func TestDevStorageEndpoints(t *testing.T) {
	testData := []string{"127.0.0.1:2379"}
	configData := GetStorageEndpoints()
	utils.AssertEqual(t, len(configData), len(testData), "GetStorageEndpoints() failed")

	for i := 0; i < len(testData); i++ {
		utils.AssertEqual(t, configData[i], testData[i], "GetStorageEndpoints() failed")
	}
}

func TestDevStorageRoot(t *testing.T) {
	utils.AssertEqual(t, GetStorageRoot(), "/hydra-dev", "GetStorageRoot() failed")
}
