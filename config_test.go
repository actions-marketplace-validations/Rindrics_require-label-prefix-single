package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return a new Config", func(t *testing.T) {
		// Setup environmental valiables
		//   wille be cleaned up after the test
		os.Setenv("REPOSITORY_FULL_NAME", "octocat/sandbox")
		os.Setenv("TOKEN", "abcdefghijklmnopqrstuvwxyz1234567890")
		os.Setenv("ADD_LABEL", "true")
		os.Setenv("DEFAULT_LABEL", "MyGreatLabel")
		os.Setenv("PREFIX", "foo")
		os.Setenv("SEPARATOR", "/")
		os.Setenv("COMMENT", "This is a test comment.")

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
		os.Unsetenv("REPOSITORY_FULL_NAME")
		os.Unsetenv("TOKEN")
		os.Unsetenv("ADD_LABEL")
		os.Unsetenv("DEFAULT_LABEL")
		os.Unsetenv("PREFIX")
		os.Unsetenv("SEPARATOR")
		os.Unsetenv("COMMENT")
	})
}
