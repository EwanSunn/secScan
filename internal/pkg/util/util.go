package util

import (
	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/sirupsen/logrus"
	"gopkg.in/cheggaaa/pb.v2"
	"net"
	"sync"
)

var (
	AliveAddr []model.IpAddr
	mutex     sync.Mutex
)

func init() {
	AliveAddr = make([]model.IpAddr, 0)
}

func CheckAlive(ipList []model.IpAddr) []model.IpAddr {
	logrus.Infoln("checking ip active")
	vars.ProcessBarActive = pb.StartNew(len(ipList))
	vars.ProcessBarActive.SetTemplate(`{{ rndcolor "Checking progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor}}  {{rtime . | rndcolor }}`)

	var wg sync.WaitGroup
	wg.Add(len(ipList))

	for _, addr := range ipList {
		go func(addr model.IpAddr) {
			defer wg.Done()
			SaveAddr(check(addr))
		}(addr)
	}
	wg.Wait()
	vars.ProcessBarActive.Finish()

	return AliveAddr
}

func check(ipAddr model.IpAddr) (bool, model.IpAddr) {
	alive := false
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr.Ip, ipAddr.Port), vars.TimeOut)
	if err == nil {
		alive = true
	}

	vars.ProcessBarActive.Increment()
	return alive, ipAddr
}

func SaveAddr(alive bool, ipAddr model.IpAddr) {
	if alive {
		mutex.Lock()
		AliveAddr = append(AliveAddr, ipAddr)
		mutex.Unlock()
	}
}