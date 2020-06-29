package framework

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *Logger

type Logger struct {
	fields logrus.Fields
	log    *logrus.Logger
}

func InitLogger(logFile, serviceName string) {
	var log = logrus.New()
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.WarnLevel)
	log.SetOutput(f)

	fld := logrus.Fields{
		"Service": serviceName,
	}

	logger = &Logger{fld, log}
}

func (l *Logger) Debug(args ...interface{}) {
	l.log.WithFields(l.fields).Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log.WithFields(l.fields).Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log.WithFields(l.fields).Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.log.WithFields(l.fields).Error(args...)
}

func (l *Logger) Get() io.Writer {
	return l.log.Writer()
}

func GetLoggerInstance() *Logger {
	return logger
}
