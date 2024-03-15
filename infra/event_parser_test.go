package infra

import (
	"testing"

	"github.com/caitlinelfring/go-env-default"
	"github.com/google/go-github/github"
)

func TestParseEvent(t *testing.T) {
	title := "foobar"
	event := github.IssuesEvent{
		Issue: &github.Issue{
			Title: &title,
		},
	}

	ParseEvent(&event, ParseLogLevel(env.GetDefault("LOG_LEVEL", "info")))
}
