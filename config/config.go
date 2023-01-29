// Package config exposes functions to fetch Profiles in various ways
package config

import (
	"bufio"
	"code-profiles/utils"
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles    []Profile
	BaseProfile string `yaml:"base-profile"`
}
type Profile struct {
	Name       string     `yaml:"name"`
	Path       string     `yaml:"profile-path"`
	Extensions []string   `yaml:"extensions,flow"`
	DependsOn  [][]string `yaml:"depends-on,flow"`
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

// GetProfile returns either the named profile or loads the .code-profile value & tries to return this profile.
// returns an error if profile can't be found.
func GetProfile(name string) (Profile, error) {
	if name == "" {
		return getProfileFromFile()
	} else {
		return getProfileByName(name)
	}
}

// getProfileByName returns a profile from loaded config given its name.
func getProfileByName(name string) (Profile, error) {
	for _, p := range instance().Profiles {
		if p.Name == name {
			return p, nil
		}
	}
	return Profile{}, errors.New("no profile with name '" + name + "' found.")
}

// GetProfileFromFile fetches the project from the '.code-profile' file in the current working dir.
// '.code-profile' file should only contain the name of the profile on the first line.
func getProfileFromFile() (Profile, error) {
	file, err := os.Open(".code-profile")
	defer file.Close()
	if err != nil {
		return Profile{}, errors.New("cannot find file '.code-profile'")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err != nil {
		return Profile{}, err
	}
	return getProfileByName(string(scanner.Text()))
}
