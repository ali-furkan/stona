package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ali-furkqn/stona/internal/pkg/check"
	"gopkg.in/yaml.v2"
)

// ConfigLoad loads file which entered path parameter as configuration
// If path is nil, It defines default configüration
func configLoad(path string) *ConfigStruct {
	config := ConfigStruct{}

	// Default Configurations
	config.Connection.Address = "127.0.0.1"
	config.Connection.Port = 7050
	config.Connection.Timeout = 15

	config.RootPath = "/"

	if path == "" {
		return &config
	}

	// Get Config Path
	base, err := os.Getwd()
	check.Panic(err)

	configFilePath, err := filepath.Abs(fmt.Sprintf("%s/%s", base, path))
	check.Panic(err)

	// Parse Config File
	dat, err := ioutil.ReadFile(configFilePath)
	check.Panic(err)

	err = yaml.Unmarshal(dat, &config)
	check.Panic(err)

	// Handle Config File
	if config.Buckets != nil {
		for _, bucket := range config.Buckets {
			InitBucketConfig(bucket)
		}
	}

	return &config
}
