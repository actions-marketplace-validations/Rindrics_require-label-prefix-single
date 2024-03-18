package application

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return a new Config", func(t *testing.T) {
		// Setup environmental valiables
		//   wille be cleaned up after the test
		os.Setenv("GITHUB_REPOSITORY", "octocat/sandbox")
		os.Setenv("INPUT_TOKEN", "abcdefghijklmnopqrstuvwxyz1234567890")
		os.Setenv("INPUT_ADD_LABEL", "true")
		os.Setenv("INPUT_DEFAULT_LABEL", "MyGreatLabel")
		os.Setenv("INPUT_LABEL_PREFIX", "foo")
		os.Setenv("INPUT_LABEL_SEPARATOR", "/")
		os.Setenv("INPUT_COMMENT", "This is a test comment.")

		config := NewConfig()

		want := &Config{
			Owner:        "octocat",
			Repository:   "sandbox",
			Token:        "abcdefghijklmnopqrstuvwxyz1234567890",
			AddLabel:     true,
			DefaultLabel: "MyGreatLabel",
			Prefix:       "foo",
			Separator:    "/",
			Comment:      "This is a test comment.",
		}

		if !reflect.DeepEqual(config, want) {
			t.Errorf("got %v, want %v", config, want)
		}

		// Clean up environmental variables
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_TOKEN")
		os.Unsetenv("INPUT_ADD_LABEL")
		os.Unsetenv("INPUT_DEFAULT_LABEL")
		os.Unsetenv("INPUT_LABEL_PREFIX")
		os.Unsetenv("INPUT_LABEL_SEPARATOR")
		os.Unsetenv("INPUT_COMMENT")
	})
}
