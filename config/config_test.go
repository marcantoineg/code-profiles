package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	baseProfileBody = `
profiles:
  - name: base
    extensions:
      - some.very.real.ext.id
`
	multipleProfilesBody = `
profiles:
  - name: base
    extensions:
      - some.very.real.ext.id

  - name: base2
    extensions:
      - some.very.real.ext.id
`

	invalidProfiles = `
profiles:
    - name: base
	extensions:
`

	testRuns = []struct {
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

			&Profile{Name: "base", Extensions: []string{"some.very.real.ext.id"}},
			false,
			false,
		},
		{
			"test with multiple valid profiles on disk and valid name",
			multipleProfilesBody,
			"base2",

			&Profile{Name: "base2", Extensions: []string{"some.very.real.ext.id"}},
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
	os.Remove("code-profiles.yml")
	os.Remove(".code-profile")

	code := m.Run()

	// teardown
	os.Remove("code-profiles.yml")
	os.Remove(".code-profile")

	os.Exit(code)
}

func Test_GetProfile(t *testing.T) {
	for _, tr := range testRuns {
		t.Run(tr.testName, func(t *testing.T) {
			os.Remove("code-profiles.yml")

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

func Test_GetProfileFromFile(t *testing.T) {
	for _, tr := range testRuns {
		t.Run(tr.testName, func(t *testing.T) {
			os.Remove("code-profiles.yml")
			os.ReadDir(".code-profile")

			saveStringToFile("code-profiles.yml", tr.initialDiskData)
			saveStringToFile(".code-profile", tr.profileNameToFetch)

			if tr.expectsPanic {
				assert.Panics(t, func() { GetProfileFromFile() })
			} else {
				actual, err := GetProfileFromFile()

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
