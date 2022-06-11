package cmd

import (
	"github.com/EwanSunn/secScan/internal/config"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/EwanSunn/secScan/internal/pkg/util"
	scanner "github.com/EwanSunn/secScan/internal/scan/port-scan"
	"github.com/desertbit/grumble"
	"github.com/sirupsen/logrus"
	"net"
)

var PortScan = &grumble.Command{
	Name:  "portScan",
	Help:  "Perform a port scan for the target",
	Usage: "portScan [-i 127.0.0.1(or 127.0.0.1/24)] [-f ./ip.txt] -p 22,23-24 -c 1000",
	Run:   runPortScan,
	Flags: func(f *grumble.Flags) {
		f.Bool("d", "debug", false, "debug mod")
		f.String("i", "ip", "", "ip list")
		f.String("f", "file", "", "ip list file")
		f.String("p", "port", "", "port list")
		f.Int("c", "thread", 1000, "thread num")
	},
}

func runPortScan(ctx *grumble.Context) (err error) {
	var (
		ips []net.IP
	)
	if ctx.Flags.String("ip") != "" {
		ipList := ctx.Flags.String("ip")
		ips, err = util.GetIpList(ipList)
	} else if ctx.Flags.String("ip") == "" && ctx.Flags.String("file") != "" {
		ipFile := ctx.Flags.String("file")
		ips, err = util.ReadIpList(ipFile)
	} else {
		config.Config.Log.Error("Invalid arguments.")
	}
	if ctx.Flags.Bool("debug") != false {
		vars.DebugMode = true
		config.Config.Log.Level = logrus.DebugLevel
	}

	portList := ctx.Flags.String("port")
	threadNum := ctx.Flags.Int("thread")
	ports, err := util.GetPorts(portList)
	_ = err

	tasks, _ := scanner.GenerateTask(ips, ports)
	scanner.RunTask(tasks, threadNum)
	scanner.PrintResult()
	return
}
