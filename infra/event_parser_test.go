package infra

import (
	"testing"

	"github.com/google/go-github/github"
)

func TestParseEvent(t *testing.T) {
	title := "foobar"
	number := 123
	label_names := []string{"foo", "bar", "baz"}
	event := github.IssuesEvent{
		Issue: &github.Issue{
			Title:  &title,
			Number: &number,
			Labels: []github.Label{
				github.Label{Name: &label_names[0]},
				github.Label{Name: &label_names[1]},
				github.Label{Name: &label_names[2]},
			},
		},
	}

	ParseEvent(&event, ParseLogLevel())
}
