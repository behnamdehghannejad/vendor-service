package log

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func Initialize() error {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return err
	}
	logger = logrus.New()
	logger.SetOutput(file)

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(logrus.DebugLevel)
	return nil
}

func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

func Info(msg string) {
	logger.Info(msg)
}

func Warningf(format string, args ...any) {
	logger.Warnf(format, args...)
}

func Warning(msg string) {
	logger.Warn(msg)
}

func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

func Error(msg string) {
	logger.Error(msg)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}

func Fatal(msg string) {
	logger.Fatal(msg)
}
