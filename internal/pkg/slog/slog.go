package slog

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true                    // 显示完整时间
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	customFormatter.DisableColors = false                   // 禁止颜色显示
	logger = logrus.New()
	logger.SetFormatter(customFormatter)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func SetDebug(debug bool) {
	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func Info(args ...interface{}) {
	logger.Info(args)
}
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args)
}
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
