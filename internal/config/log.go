package config

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

func initLogConfig() {
	level, _ := logger.ParseLevel(mustGetString("LOG_LEVEL"))

	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetReportCaller(true)
}
