package config

import (
	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	return log
}
