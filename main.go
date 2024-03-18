package main

import (
	"flag"
	"os"

	"github.com/Rindrics/require-label-prefix-single/application"
	"github.com/Rindrics/require-label-prefix-single/domain"
	"github.com/Rindrics/require-label-prefix-single/infra"
)

func main() {
	logLevelFlag := flag.String("log-level", "info", "Set the logging level (debug, info, warn, error)")
	flag.Parse()

	logger := infra.ParseLogLevel(*logLevelFlag)

	logger.Info("Loading webhook event")
	event, err := infra.LoadEventFromEnv()
	if err != nil {
		logger.Error("Failed to load event from environment", "error", err)
		return
	}

	logger.Info("Parsing webhook event")
	eventInfo := infra.ParseEvent(event, logger)

	config := application.NewConfig()

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

	app := application.New(eventInfo, client, *config, logger)
	err = app.Run()
	if err != nil {
		logger.Error("Error running application", "error", err)
		os.Exit(1)
	}
}
