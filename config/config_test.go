package config

func Test_GetProfile(t *testing.T) {
	testRuns := []struct {
		testName string
		initialDiskData string

		expectedData config.Profile
		expectedError error
	} {
		{
			"test with valid project on disk",
			```
			profiles:
			  - name: base
			    extensions:
			      - some.very.real.ext.id
			```,

			project{Name: "base", extensions: {"some.very.real.ext.id"}},
			nil,
		}
	}

	
}