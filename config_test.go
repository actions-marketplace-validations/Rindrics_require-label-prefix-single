package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return a new Config", func(t *testing.T) {
		os.Args = []string{"cmd",
			"-repository", "octocat/sandbox",
			"-token", "abcdefghijklmnopqrstuvwxyz1234567890",
			"-add-label", "true",
			"-default-label", "My default label",
			"-prefix", "foo",
			"-separator", "/",
		}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		got := NewConfig()
		want := &Config{
			Repository:   "octocat/sandbox",
			Token:        "abcdefghijklmnopqrstuvwxyz1234567890",
			AddLabel:     "true",
			DefaultLabel: "My default label",
			Prefix:       "foo",
			Separator:    "/",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
