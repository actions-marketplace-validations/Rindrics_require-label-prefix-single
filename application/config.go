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
	repositoryFullName := os.Getenv("GITHUB_REPOSITORY")
	token := os.Getenv("INPUT_TOKEN")
	addLabel := env.GetBoolDefault("INPUT_ADD_LABEL", true)
	defaultLabel := env.GetDefault("INPUT_DEFAULT_LABEL", "label-required")
	prefix := env.GetDefault("INPUT_LABEL_PREFIX", "prefix")
	separator := env.GetDefault("INPUT_LABEL_SEPARATOR", "/")
	comment := env.GetDefault("INPUT_COMMENT", "Add a label with a prefix.")

	parts := strings.Split(repositoryFullName, "/")
	if len(parts) != 2 {
		log.Fatalf("'repository_full_name' must be in 'owner/repo' format, got: %s", repositoryFullName)
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
