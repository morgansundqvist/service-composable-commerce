package adapters

import (
	"github.com/morgansundqvist/service-composable-commerce/internal/ports"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() ports.Logger {
	return &LogrusLogger{
		logger: logrus.New(),
	}
}

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(msg)
}

func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Error(msg)
}

func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Fatal(msg)
}
