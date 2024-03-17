package main

import (
	"log"
	"os"
	"strings"

	"github.com/caitlinelfring/go-env-default"
)

type Config struct {
	Owner        string
	Repository   string
	Token        string
	AddLabel     bool
	DefaultLabel string
	Prefix       string
	Separator    string
	Comment      string
}

func NewConfig() *Config {
	repositoryFullName := os.Getenv("REPOSITORY_FULL_NAME")
	token := os.Getenv("TOKEN")
	addLabel := env.GetBoolDefault("ADD_LABEL", true)
	defaultLabel := env.GetDefault("DEFAULT_LABEL", "label-required")
	prefix := env.GetDefault("PREFIX", "prefix")
	separator := env.GetDefault("SEPARATOR", "/")
	comment := env.GetDefault("COMMENT", "Add a label with a prefix.")

	parts := strings.Split(repositoryFullName, "/")
	if len(parts) != 2 {
		log.Fatalf("REPOSITORY_FULL_NAME must be in 'owner/repo' format, got: %s", repositoryFullName)
	}

	return &Config{
		Owner:        parts[0],
		Repository:   parts[1],
		Token:        token,
		AddLabel:     addLabel,
		DefaultLabel: defaultLabel,
		Prefix:       prefix,
		Separator:    separator,
		Comment:      comment,
	}
}
