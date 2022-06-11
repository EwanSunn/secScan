package config

import (
	"github.com/EwanSunn/secScan/internal/pkg/logger"
	"github.com/sirupsen/logrus"
	"time"
)

var Config config

//CONFIG
type config struct {
	ThreadNum int
	Timeout   time.Duration
	Log       *logrus.Logger
}

func InitConfig() {
	Config = config{}

	Config.ThreadNum = 5000
	Config.Timeout = 3 * time.Second
	Config.Log = logger.InitLog()
}
