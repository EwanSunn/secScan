package task

import (
	"fmt"
	"github.com/EwanSunn/secScan/internal/config"
	"github.com/EwanSunn/secScan/internal/crack/plugin"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/EwanSunn/secScan/internal/pkg/util/hash"
	"gopkg.in/cheggaaa/pb.v2"
	"runtime"
	"strings"
	"sync"
	"time"
)

func GenerateTask(ipList []model.IpAddr, users []string, passwords []string) (tasks []model.Service, taskNum int) {
	tasks = make([]model.Service, 0)

	for _, user := range users {
		for _, password := range passwords {
			for _, addr := range ipList {
				service := model.Service{Ip: addr.Ip, Port: addr.Port, Protocol: addr.Protocol, Username: user, Password: password}
				tasks = append(tasks, service)
			}
		}
	}

	return tasks, len(tasks)
}

func RunTask(tasks []model.Service) {
	totalTask := len(tasks)
	vars.ProgressBarPassword = pb.StartNew(totalTask)
	vars.ProgressBarPassword.SetTemplate(`{{ rndcolor "Scanning progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor }} {{rtime . | rndcolor}} `)

	wg := &sync.WaitGroup{}

	// 创建一个buffer为vars.threadNum的channel
	taskChan := make(chan model.Service, vars.ScanNum)

	// 创建vars.ThreadNum个协程
	for i := 0; i < vars.ScanNum; i++ {
		go crackPassword(taskChan, wg)
	}

	// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	waitTimeout(wg, vars.TimeOut)

	// 内存中的扫描结果落盘，并导出为一个txt文件。
	{
		_ = model.SaveResultToFile()
		model.ResultTotal()
		_ = model.DumpToFile(vars.ResultFile)
	}

}

// 每个协程都从channel中读取数据后开始扫描并保存
func crackPassword(taskChan chan model.Service, wg *sync.WaitGroup) {
	for task := range taskChan {
		vars.ProgressBarPassword.Increment()

		if vars.DebugMode {
			config.Config.Log.Debugf("checking: Ip: %v, Port: %v, [%v], UserName: %v, Password: %v, goroutineNum: %v", task.Ip, task.Port,
				task.Protocol, task.Username, task.Password, runtime.NumGoroutine())
		}

		var k string
		protocol := strings.ToUpper(task.Protocol)

		if protocol == "REDIS" {
			k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Username)
		}

		h := hash.MakeTaskHash(k)
		if hash.CheckTaskHash(h) {
			wg.Done()
			continue
		}

		fn := plugin.ScanFuncMap[protocol]
		model.SaveResult(fn(task))
		wg.Done()
	}
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
