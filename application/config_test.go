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
		os.Setenv("repository_full_name", "octocat/sandbox")
		os.Setenv("token", "abcdefghijklmnopqrstuvwxyz1234567890")
		os.Setenv("add_label", "true")
		os.Setenv("default_label", "MyGreatLabel")
		os.Setenv("label_prefix", "foo")
		os.Setenv("label_separator", "/")
		os.Setenv("comment", "This is a test comment.")

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
		os.Unsetenv("repository_full_name")
		os.Unsetenv("token")
		os.Unsetenv("add_label")
		os.Unsetenv("default_label")
		os.Unsetenv("label_prefix")
		os.Unsetenv("label_separator")
		os.Unsetenv("comment")
	})
}
