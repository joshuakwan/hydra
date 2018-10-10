package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

// Configuration wraps the configuration of the application
type Configuration struct {
	StorageType      string
	StorageEndpoints []string
	StorageRoot      string
}

var globalConfig *Configuration

// GetStorageType returns the storage type
func GetStorageType() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.StorageType
}

// GetStorageEndpoints returns the storage cluster's endpoints
func GetStorageEndpoints() []string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.StorageEndpoints
}

// GetStorageRoot returns the root doc/db of the application
func GetStorageRoot() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.StorageRoot
}

func loadConfiguration() *Configuration {
	env := os.Getenv("APP_ENV")
	var filepath string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	if env == "DEV" || env == "" {
		filepath = path.Dir(filename) + "/dev.json"
	}

	if env == "TEST" {
		filepath = path.Dir(filename) + "/unittest.json"
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	return &config
}
