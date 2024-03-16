package infra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"

	"github.com/Rindrics/require-label-prefix-on-closed/domain"
	"github.com/google/go-github/github"
)

func LoadEventFromEnv() (*github.IssuesEvent, error) {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_PATH environment variable not set")
	}

	data, err := ioutil.ReadFile(eventPath)
	if err != nil {
		return nil, err
	}

	var event github.IssuesEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func ParseEvent(event *github.IssuesEvent, l *slog.Logger) domain.EventInfo {
	title := event.GetIssue().GetTitle()
	number := event.GetIssue().GetNumber()

	var labels []string
	for _, label := range event.GetIssue().Labels {
		labels = append(labels, *label.Name)
	}

	l.Debug("parsed", "title", title)
	l.Debug("parsed", "number", number)
	l.Debug("parsed", "labels", labels)

	return domain.EventInfo{number, labels}
}
