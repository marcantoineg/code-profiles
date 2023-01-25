package config

import (
	"bufio"
	"code-profiles/utils"
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles    []Profile
	BaseProfile string `yaml:"base-profile"`
}
type Profile struct {
	Name       string   `yaml:"name"`
	Path       string   `yaml:"profile-path"`
	Extensions []string `yaml:"extensions,flow"`
}

var _instance *config = nil

// instance creates and/or gets the current config singleton.
func instance() *config {
	if _instance == nil {
		_instance = &config{}
		_instance.init_config()
	}
	return _instance
}

// init_config reads the config on disk and loads it in memory,
func (_ config) init_config() {
	data, err := os.ReadFile("./code-profiles.yml")
	utils.Check(err)

	_instance = nil
	err = yaml.Unmarshal(data, &_instance)
	utils.Check(err)
}

// GetProfile returns the a profile given its name.
func GetProfile(name string) (Profile, error) {
	for _, p := range instance().Profiles {
		if p.Name == name {
			return p, nil
		}
	}
	return Profile{}, errors.New("no profile with name " + name)
}

// GetProfileFromFile fetches the project from the '.code-profile' file.
// '.code-profile' file should only contain the name of the profile.
func GetProfileFromFile() (Profile, error) {
	file, err := os.Open(".code-profile")
	if err != nil {
		return Profile{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err != nil {
		return Profile{}, err
	}
	return GetProfile(string(scanner.Text()))

}

// GetProfile returns the a base profile if specified.
func BaseProfile() (Profile, error) {
	return GetProfile(instance().BaseProfile)
}