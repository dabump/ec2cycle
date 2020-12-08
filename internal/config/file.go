package config

import (
	"os"
	"path/filepath"
)

type configFile struct {
	configName string
	configType string
}

func newConfigFile() *configFile {
	return &configFile{
		configName: "config",
		configType: "yaml",
	}
}

func (ac *configFile) path() string {
	return filepath.Join(ac.configName + "." + ac.configType)
}

func (ac *configFile) exists() bool {
	_, err := os.Stat(ac.path())
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func (ac *configFile) create() error {
	_, err := os.Create(ac.path())
	return err
}
