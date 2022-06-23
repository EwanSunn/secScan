package config

import (
	"time"
)

var Config config

//CONFIG
type config struct {
	ThreadNum int
	Timeout   time.Duration
}

func InitConfig() {
	Config = config{}

	Config.ThreadNum = 5000
	Config.Timeout = 3 * time.Second
}
