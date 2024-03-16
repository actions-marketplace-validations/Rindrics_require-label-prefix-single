package main

import (
	"flag"
	"strings"
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
	config := &Config{}
	var repositoryFullName string

	flag.StringVar(&repositoryFullName, "repository-full-name", "", "Name of repository in 'owner/repo' format")
	flag.StringVar(&config.Token, "token", "", "GitHub access token")
	flag.BoolVar(&config.AddLabel, "add-label", false, "Whether add label or not if the issue did not have label with required prefix")
	flag.StringVar(&config.DefaultLabel, "default-label", "", "A label to be used when 'add-label=true'")
	flag.StringVar(&config.Prefix, "prefix", "", "Required label prefix to look for")
	flag.StringVar(&config.Separator, "separator", "", "Character which separates prefix and label body")
	flag.StringVar(&config.Comment, "comment", "", "Comment to post when 'add-label=false'")

	flag.Parse()

	if repositoryFullName != "" {
		parts := strings.Split(repositoryFullName, "/")
		if len(parts) != 2 {
			panic("repository-full-name must be in 'owner/repo' format")
		}
		config.Owner, config.Repository = parts[0], parts[1]

	}

	return config
}
