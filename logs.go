package torchlight

import "github.com/sirupsen/logrus"

func GetLogger(verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}
