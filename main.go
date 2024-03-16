package main

import (
	"flag"

	"github.com/Rindrics/require-label-prefix-on-closed/infra"
)

func main() {
	logLevelFlag := flag.String("log-level", "info", "Set the logging level (debug, info, warn, error)")
	flag.Parse()

	logger := infra.ParseLogLevel(*logLevelFlag)

	e, err := infra.LoadEventFromEnv()
	if err != nil {
		logger.Error("Failed to load event from environment", "error", err)
		return
	}
	eventInfo := infra.ParseEvent(e, logger)

	logger.Info("Parsing webhook event")
	logger.Debug("event info", "number", eventInfo.Number, "labels", eventInfo.Labels)
}
