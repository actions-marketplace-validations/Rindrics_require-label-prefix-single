package main

import "flag"

type Config struct {
	Repository   string
	Token        string
	AddLabel     string
	DefaultLabel string
	Prefix       string
	Separator    string
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Repository, "repository", "", "Name of GitHub repository")
	flag.StringVar(&config.Token, "token", "", "GitHub access token")
	flag.StringVar(&config.AddLabel, "add-label", "", "Whether add label or not if the issue did not have label with required prefix")
	flag.StringVar(&config.DefaultLabel, "default-label", "", "Default label to be used when 'add-label=true'")
	flag.StringVar(&config.Prefix, "prefix", "", "Required label prefix to look for")
	flag.StringVar(&config.Separator, "separator", "", "Character which separates prefix and label body")

	flag.Parse()

	return config
}
