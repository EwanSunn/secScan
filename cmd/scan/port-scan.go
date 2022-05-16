package cmd

import (
	"github.com/EwanSunn/secScan/internal/pkg/util"
	scanner "github.com/EwanSunn/secScan/internal/scan/port-scan"
	"github.com/desertbit/grumble"
)

var PortScan = &grumble.Command{
	Name: "portScan",
	Help: "Perform a port scan for the target",
	Run:  runPortScan,
	Flags: func(f *grumble.Flags) {
		f.String("i","ip","","ip list or ip file")
		f.String("p", "port", "", "port list")
		f.Int("n","thread",1000,"thread num")
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
