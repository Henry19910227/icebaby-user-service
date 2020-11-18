package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GPLogger ...
type GPLogger struct {
	RunMode string
	print   *logrus.Logger
	write   *logrus.Logger
}

// NewGPLogger ...
func NewGPLogger(setting LogSetting) (*GPLogger, error) {
	writeLog, err := newWriteLogger(setting)
	if err != nil {
		return nil, err
	}
	printLog := newPrintLogger()
	runMode := setting.GetRunMode()

	return &GPLogger{runMode, printLog, writeLog}, nil
}

func newPrintLogger() *logrus.Logger {
	printLog := logrus.New()
	printLog.SetFormatter(&logrus.JSONFormatter{})
	printLog.SetOutput(os.Stdout)
	return printLog
}

func newWriteLogger(setting LogSetting) (*logrus.Logger, error) {
	file, err := os.OpenFile(setting.GetLogFilePath()+"/"+setting.GetLogFileName()+"."+setting.GetLogFileExt(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	writeLog := logrus.New()
	writeLog.SetFormatter(&logrus.JSONFormatter{})
	writeLog.SetOutput(file)
	return writeLog, nil
}

// Trace implement Logger interface
func (logger *GPLogger) Trace(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Trace(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Trace(msg)
	}

}

// Debug implement Logger interface
func (logger *GPLogger) Debug(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Debug(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Debug(msg)
	}
}

// Info implement Logger interface
func (logger *GPLogger) Info(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Info(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Info(msg)
	}
}

// Warn implement Logger interface
func (logger *GPLogger) Warn(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Warn(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Warn(msg)
	}
}

// Error implement Logger interface
func (logger *GPLogger) Error(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Error(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Error(msg)
	}
}

// Fatal implement Logger interface
func (logger *GPLogger) Fatal(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Fatal(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Fatal(msg)
	}
}

// Panic implement Logger interface
func (logger *GPLogger) Panic(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Panic(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Panic(msg)
	}
}
