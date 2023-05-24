package logging

import "github.com/sirupsen/logrus"

type AppLogger = logrus.FieldLogger

type Config struct {
	LogLevel int
}

func New(config Config) AppLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.Level(config.LogLevel))
	return logger
}
