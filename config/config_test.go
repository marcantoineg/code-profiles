package config

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	baseProfileBody = `
profiles:
  - name: base
    profile-path: ./base
    extensions:
      - some.very.real.ext.id
`
	multipleProfilesBody = `
profiles:
  - name: base
    profile-path: ./base
    extensions:
      - some.very.real.ext.id

  - name: base2
    profile-path: ./base2
    extensions:
      - some.very.real.ext.id
`

	invalidProfiles = `
profiles:
    - name: base
	extensions:
`

	sharedTestRuns = []struct {
		testName           string
		initialDiskData    string
		profileNameToFetch string

		expectedData *Profile
		expectsError bool
		expectsPanic bool
	}{
		{
			"test with single valid profile on disk and valid id",
			baseProfileBody,
			"base",

			&Profile{Name: "base", ProfilePath: "./base", Extensions: []string{"some.very.real.ext.id"}},
			false,
			false,
		},
		{
			"test with multiple valid profiles on disk and valid name",
			multipleProfilesBody,
			"base2",

			&Profile{Name: "base2", ProfilePath: "./base2", Extensions: []string{"some.very.real.ext.id"}},
			false,
			false,
		},
		{
			"test with invalid config file panics",
			invalidProfiles,
			"base2",

			nil,
			false,
			true,
		},
		{
			"test with multiple valid profiles and invalid name",
			multipleProfilesBody,
			"base3",

			&Profile{},
			true,
			false,
		},
	}
)

func TestMain(m *testing.M) {
	// setup
	_configPath = "./code-profiles.yml"
	os.Remove("code-profiles.yml")
	os.Remove(".code-profile")

	code := m.Run()

	// teardown
	os.Remove("code-profiles.yml")
	os.Remove(".code-profile")

	os.Exit(code)
}

func Test_GetProfile(t *testing.T) {
	runs := []struct {
		testName            string
		profilesFileData    string
		ProfileIdFileData   string
		deleteProfileIdFile bool
		profileIdParam      string

		expectedData  *Profile
		expectedError error
	}{
		{
			"test with single valid profile on disk and valid id",
			baseProfileBody,
			"",
			false,
			"base",

			&Profile{Name: "base", ProfilePath: "./base", Extensions: []string{"some.very.real.ext.id"}},
			nil,
		},
		{
			"test with multiple valid profile on disk and valid id",
			multipleProfilesBody,
			"",
			false,
			"base",

			&Profile{Name: "base", ProfilePath: "./base", Extensions: []string{"some.very.real.ext.id"}},
			nil,
		},
		{
			"test with single valid profile on disk, with id on disk and empty id as param",
			baseProfileBody,
			"base",
			false,
			"",

			&Profile{Name: "base", ProfilePath: "./base", Extensions: []string{"some.very.real.ext.id"}},
			nil,
		},
		{
			"test with single valid profile on disk, with invalid id on disk and empty id as param",
			baseProfileBody,
			"base2",
			false,
			"",

			&Profile{},
			errors.New("no profile with name 'base2' found."),
		},
		{
			"test with single valid profile on disk, with inexistant id on disk and invalid id as param",
			baseProfileBody,
			"",
			true,
			"base2",

			&Profile{},
			errors.New("no profile with name 'base2' found."),
		},
		{
			"test with single valid profile on disk, with inexistant id on disk and empty id as param",
			baseProfileBody,
			"",
			true,
			"",

			&Profile{},
			errors.New("cannot find file '.code-profile'"),
		},
		{
			"test with single valid profile on disk, with invalid id on disk and empty id as param",
			baseProfileBody,
			"base2",
			false,
			"",

			&Profile{},
			errors.New("no profile with name 'base2' found."),
		},
	}

	for _, tr := range runs {
		t.Run(tr.testName, func(t *testing.T) {
			os.Remove("code-profiles.yml")
			os.Remove(".code-profile")
			_configPath = "./code-profiles.yml"

			saveStringToFile("code-profiles.yml", tr.profilesFileData)

			if !tr.deleteProfileIdFile {
				saveStringToFile(".code-profile", tr.ProfileIdFileData)
			}

			actual, err := GetProfile(tr.profileIdParam)

			assert.Equal(t, tr.expectedError, err)

			if tr.expectedData != nil {
				assert.Equal(t, *tr.expectedData, actual)
			} else {
				assert.Equal(t, Profile{}, actual)
			}
		})
	}
}

func Test_SetConfigPath(t *testing.T) {
	t.Run("initial config name is set corretly", func(t *testing.T) {
		assert.Equal(t, "./code-profiles.yml", _configPath)
	})

	_configPath = "./code-profiles.yml"
	t.Run("setting the config path invalidates the _instance", func(t *testing.T) {
		SetConfigPath("some-config-path")
		assert.Nil(t, _instance)
		assert.Equal(t, "some-config-path", _configPath)
	})

	_configPath = "./code-profiles.yml"
	t.Run("setting the config path to an empty value doesn't do anything", func(t *testing.T) {
		assert.Nil(t, _instance)
		instance()
		assert.NotNil(t, _instance)

		SetConfigPath("")
		assert.NotNil(t, _instance)
		assert.Equal(t, "./code-profiles.yml", _configPath)
	})
}

func Test_getProfileByName(t *testing.T) {
	for _, tr := range sharedTestRuns {
		t.Run(tr.testName, func(t *testing.T) {
			os.Remove("code-profiles.yml")
			_configPath = "./code-profiles.yml"

			saveStringToFile("code-profiles.yml", tr.initialDiskData)

			if tr.expectsPanic {
				assert.Panics(t, func() { GetProfile(tr.profileNameToFetch) })
			} else {
				actual, err := GetProfile(tr.profileNameToFetch)

				if tr.expectsError {
					assert.NotNil(t, err)
				}

				if tr.expectedData != nil {
					assert.Equal(t, *tr.expectedData, actual)
				} else {
					assert.Equal(t, Profile{}, actual)
				}
			}
		})
	}
}

func Test_getProfileFromFile(t *testing.T) {
	for _, tr := range sharedTestRuns {
		t.Run(tr.testName, func(t *testing.T) {
			os.Remove("code-profiles.yml")
			os.ReadDir(".code-profile")
			_configPath = "./code-profiles.yml"

			saveStringToFile("code-profiles.yml", tr.initialDiskData)
			saveStringToFile(".code-profile", tr.profileNameToFetch)

			if tr.expectsPanic {
				assert.Panics(t, func() { getProfileFromFile() })
			} else {
				actual, err := getProfileFromFile()

				if tr.expectsError {
					assert.NotNil(t, err)
				}

				if tr.expectedData != nil {
					assert.Equal(t, *tr.expectedData, actual)
				} else {
					assert.Equal(t, Profile{}, actual)
				}
			}
		})
	}
}

func saveStringToFile(filePath, data string) error {
	return ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
}
