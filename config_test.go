package main

import (
	"flag"
	"log"
	"reflect"
	"strings"
	"testing"
)

func defineStringFlag(fs *flag.FlagSet, p *string, name string, value string, usage string) {
	fs.StringVar(p, name, value, usage)
}

func defineBoolFlag(fs *flag.FlagSet, p *bool, name string, value bool, usage string) {
	fs.BoolVar(p, name, value, usage)
}

func TestNewConfig(t *testing.T) {
	t.Run("should return a new Config", func(t *testing.T) {
		testFlagSet := flag.NewFlagSet("test", flag.ExitOnError)

		var repositoryFullName, token, defaultLabel, prefix, separator, comment string
		var addLabel bool

		defineStringFlag(testFlagSet, &repositoryFullName, "repository-full-name", "", "Name of repository in 'owner/repo' format")
		defineStringFlag(testFlagSet, &token, "token", "", "GitHub access token")
		defineBoolFlag(testFlagSet, &addLabel, "add-label", true, "Whether add label or not if the issue did not have label with required prefix")
		defineStringFlag(testFlagSet, &defaultLabel, "default-label", "", "A label to be used when add-label=true")
		defineStringFlag(testFlagSet, &prefix, "prefix", "", "Required label prefix to look for")
		defineStringFlag(testFlagSet, &separator, "separator", "/", "Character which separates prefix and label body")
		defineStringFlag(testFlagSet, &comment, "comment", "Default label is added", "Comment to post when 'add-label=false'")

		testArgs := []string{"cmd",
			"-default-label", "MyGreatLabel",
			"-repository-full-name", "octocat/sandbox",
			"-prefix", "foo",
			"-separator", "/",
			"-comment", "This is a test comment.",
			"-token", "abcdefghijklmnopqrstuvwxyz1234567890",
			"-add-label", "true",
		}

		log.Printf("defaultLabel-----------: %s", comment)

		err := testFlagSet.Parse(testArgs[1:])
		if err != nil {
			t.Fatalf("Failed to parse flags: %v", err)
		}

		log.Printf("defaultLabel-------------: %s", comment)

		config := &Config{
			Owner:        strings.Split(repositoryFullName, "/")[0],
			Repository:   strings.Split(repositoryFullName, "/")[1],
			Token:        token,
			AddLabel:     addLabel,
			DefaultLabel: defaultLabel,
			Prefix:       prefix,
			Separator:    separator,
			Comment:      comment,
		}

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
	})
}
