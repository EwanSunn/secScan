package logger

import (
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/sirupsen/logrus"
)

func InitLog() *logrus.Logger {
	var logger = logrus.New()
	if vars.DebugMode == true {
		logger.Level = logrus.DebugLevel
	}
	logger.Level = logrus.InfoLevel
	return logger
}
