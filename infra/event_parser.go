package infra

import (
	"log/slog"

	"github.com/google/go-github/github"
)

func ParseEvent(event *github.IssuesEvent, l *slog.Logger) {
	title := event.GetIssue().GetTitle()

	l.Debug("parsed", "title", title)
}
