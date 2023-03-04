package config

import (
	"bufio"
	"errors"
	"os"

	"github.com/marcantoineg/fileutil"
)

type Profile struct {
	Name       string     `yaml:"name"`
	Extensions []string   `yaml:"extensions,flow"`
	DependsOn  [][]string `yaml:"depends-on,flow"`
	path       string     `yaml:"profile-path"`
}

// Path returns the absolute path of the profile.
func (p Profile) Path() string {
	return fileutil.ReplaceTilde(p.path)
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

// GetProfileFromFile fetches the profile from the '.code-profile' file in the current working dir.
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
