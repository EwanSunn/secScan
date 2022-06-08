package cmd

import (
	"github.com/EwanSunn/secScan/internal/pkg/util"
	scanner "github.com/EwanSunn/secScan/internal/scan/port-scan"
	"github.com/desertbit/grumble"
)

var PortScan = &grumble.Command{
	Name: "portScan",
	Help: "Perform a port scan for the target",
	Usage: "portScan -i 127.0.0.1(or 127.0.0.1/24) -p 22,23-24 -c 1000",
	Run:  runPortScan,
	Flags: func(f *grumble.Flags) {
		f.String("i","ip","","ip list or ip file")
		f.String("p", "port", "", "port list")
		f.Int("c","thread",1000,"thread num")
	},
}



func runPortScan(ctx *grumble.Context) (err error) {
	ipList := ctx.Flags.String("ip")
	portList := ctx.Flags.String("port")
	threadNum := ctx.Flags.Int("thread")
	ips, err := util.GetIpList(ipList)
	ports, err := util.GetPorts(portList)
	_ = err

	tasks, _ := scanner.GenerateTask(ips, ports)
	scanner.RunTask(tasks, threadNum)
	scanner.PrintResult()
	return
}
