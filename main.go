package main

import (
	"flag"
	"os"

	"github.com/Rindrics/require-label-prefix/application"
	"github.com/Rindrics/require-label-prefix/domain"
	"github.com/Rindrics/require-label-prefix/infra"
)

func main() {
	logLevelFlag := flag.String("log-level", "info", "Set the logging level (debug, info, warn, error)")
	flag.Parse()

	logger := infra.ParseLogLevel(*logLevelFlag)

	logger.Info("Loading webhook event")
	e, err := infra.LoadEventFromEnv()
	if err != nil {
		logger.Error("Failed to load event from environment", "error", err)
		return
	}

	logger.Info("Parsing webhook event")
	eventInfo := infra.ParseEvent(e, logger)

	config := NewConfig()

	logger.Debug("event info", "number", eventInfo.Number, "labels", eventInfo.Labels)

	rl := domain.RequiredLabel{
		Prefix:    config.Prefix,
		Separator: config.Separator,
	}

	found := rl.DoExist(eventInfo.Labels)
	if found {
		logger.Info("Label found")
		os.Exit(0)
	}

	logger.Info("Label not found")
	client := infra.NewGitHubClient(config.Token)

	if config.AddLabel {
		logger.Info("Adding label")
		ac := application.AddLabelsCommand{
			Labeler: client,
			Params: application.AddLabelsParams{
				Number: eventInfo.Number,
				Labels: []string{config.DefaultLabel},
			},
			OnSuccess: application.PostCommentCommand{
				Commenter: client,
				Params: application.PostCommentParams{
					RepoInfo: domain.RepoInfo{
						Owner: config.Owner,
						Repo:  config.Repository,
					},
					Number: eventInfo.Number,
					Body:   "Label added",
				},
			},
		}
		err := ac.Execute()
		if err != nil {
			logger.Error("Failed to add label", "error", err)
			return
		}
	} else {
		logger.Info("Post comment without adding label")
		pc := application.PostCommentCommand{
			Commenter: client,
			Params: application.PostCommentParams{
				RepoInfo: domain.RepoInfo{
					Owner: config.Owner,
					Repo:  config.Repository,
				},
				Number: eventInfo.Number,
				Body:   config.Comment,
			},
			OnSuccess: &application.ExitAction{},
		}
		err := pc.Execute()
		if err != nil {
			logger.Error("Failed to post comment", "error", err)
			return
		}

	}
}
