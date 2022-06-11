package crack

import (
	"github.com/EwanSunn/secScan/internal/crack/task"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/EwanSunn/secScan/internal/pkg/util"
	"github.com/desertbit/grumble"
	"time"
)

var Crack = &grumble.Command{
	Name:  "crack",
	Help:  "start to crack weak password",
	Usage: "crack -f ip.txt -u ./dict/user.dic -p ./dict/pass.dic [-t 5] [-c 1000]",
	Run:   runCrack,
	Flags: func(f *grumble.Flags) {
		f.Int("t", "timeout", 5, "timeout")
		f.Int("c", "thread", 1000, "thread num")
		f.String("u", "user", "./dict/user.dic", "user dict")
		f.String("p", "password", "./dict/pass.dic", "password dict")
		f.String("o", "outfile", "./result/crack_result.txt", "result file")
		f.String("f", "file", "./dict/ip.txt", "ip list file")
	},
}

func runCrack(ctx *grumble.Context) (err error) {
	vars.TimeOut = time.Duration(ctx.Flags.Int("timeout")) * time.Second
	vars.ScanNum = ctx.Flags.Int("thread")
	vars.IpList = ctx.Flags.String("file")
	vars.UserDict = ctx.Flags.String("user")
	vars.PassDict = ctx.Flags.String("password")
	vars.ResultFile = ctx.Flags.String("outfile")
	vars.StartTime = time.Now()
	userDict, uErr := util.ReadUserDict(vars.UserDict)
	passDict, pErr := util.ReadPasswordDict(vars.PassDict)

	ipList := util.ReadIpPortList(vars.IpList)

	aliveIpList := util.CheckAlive(ipList)
	if uErr == nil && pErr == nil {
		tasks, _ := task.GenerateTask(aliveIpList, userDict, passDict)
		task.RunTask(tasks)
	}

	return err
}
