package config

import (
	"os"
	"path/filepath"
)

type file struct {
	configName string
	configType string
}

type configFile interface {
	path() string
	exists() bool
	create() error
}

func newConfigFile() configFile {
	return &file{
		configType: "yaml",
		configName: "config",
	}
}

func (ac *file) exists() bool {
	_, err := os.Stat(ac.path())
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func (ac *file) create() error {
	_, err := os.Create(ac.path())
	return err
}

func (ac *file) path() string {
	return filepath.Join(ac.configName + "." + ac.configType)
}
