package application

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
	repositoryFullName := os.Getenv("repository_full_name")
	token := os.Getenv("token")
	addLabel := env.GetBoolDefault("add_label", true)
	defaultLabel := env.GetDefault("default_label", "label-required")
	prefix := env.GetDefault("label_prefix", "prefix")
	separator := env.GetDefault("label_separator", "/")
	comment := env.GetDefault("comment", "Add a label with a prefix.")

	parts := strings.Split(repositoryFullName, "/")
	if len(parts) != 2 {
		log.Fatalf("repository_full_name must be in 'owner/repo' format, got: %s", repositoryFullName)
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
