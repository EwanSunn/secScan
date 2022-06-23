package main

import (
	"github.com/EwanSunn/secScan/cmd"
	"github.com/EwanSunn/secScan/internal/config"
)

func main() {
	config.InitConfig()
	cmd.Run()
}
