// Package config exposes functions to fetch Profiles in various ways
package config

import (
	"code-profiles/utils"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles    []Profile
	BaseProfile string `yaml:"base-profile"`
}

var _instance *config = nil
var _configPath string = "./code-profiles.yml"

// instance creates and/or gets the current config singleton.
func instance() *config {
	// always reloads the config from disk if running tests.
	if _instance == nil || flag.Lookup("test.v") != nil {
		_instance = &config{}
		_instance.init_config()
	}
	return _instance
}

// init_config reads the config on disk and loads it in memory,
func (_ config) init_config() {
	data, err := os.ReadFile(_configPath)
	if err != nil {
		println("cannot open file '" + _configPath + "'")
		os.Exit(0)
	}

	_instance = nil
	err = yaml.Unmarshal(data, &_instance)
	utils.Check(err)
}

// SetConfigPath sets the config path to load from & invalidates current instance.
func SetConfigPath(path string) {
	if path != "" {
		_configPath = path
		_instance = nil
	}
}
