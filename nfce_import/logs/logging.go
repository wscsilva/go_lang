package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger é uma interface para logging
type LoggerLogrus interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// NewLogger cria um novo logger
func NewLogger() LoggerLogrus {
	return &logrusLogger{}
}

type logrusLogger struct{}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l *logrusLogger) Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
