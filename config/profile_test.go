package config

import (
	"code-profiles/utils"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	runs := []struct {
		testName     string
		givenProfile Profile
		expectedPath string
	}{
		{
			"with simple profile path",
			Profile{ExtsInstallPath: "./test"},
			"./test",
		},
		{
			"with simple path starting with ~",
			Profile{ExtsInstallPath: "~"},
			"home_dir",
		},
		{
			"with more complex path starting with ~",
			Profile{ExtsInstallPath: "~/some-dir"},
			"home_dir/some-dir",
		},
		{
			"with path containing a ~",
			Profile{ExtsInstallPath: "./~"},
			"./~",
		},
	}

	for _, r := range runs {
		t.Run(r.testName, func(tt *testing.T) {
			homeDir, err := os.UserHomeDir()
			utils.Check(err)

			expectedPath := strings.Replace(r.expectedPath, "home_dir", homeDir, 1)

			assert.Equal(tt, expectedPath, r.givenProfile.Path())
		})
	}
}
