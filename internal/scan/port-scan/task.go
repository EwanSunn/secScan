package port_scan

import (
	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/EwanSunn/secScan/internal/pkg/slog"
	"gopkg.in/cheggaaa/pb.v2"
	"net"
	"os"
	"strings"
	"sync"
)

func GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	tasks := make([]map[string]int, 0)

	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String(): port}
			tasks = append(tasks, ipPort)
		}
	}

	return tasks, len(tasks)
}

func RunTask(tasks []map[string]int, threadNum int) {
	totalTask := len(tasks)
	vars.ProgressBarPort = pb.StartNew(totalTask)
	vars.ProgressBarPort.SetTemplate(`{{ rndcolor "PortScan progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor }} {{rtime . | rndcolor}} `)

	wg := &sync.WaitGroup{}

	// 创建一个buffer为vars.threadNum * 2的channel
	taskChan := make(chan map[string]int, threadNum*2)

	// 创建vars.ThreadNum个协程
	for i := 0; i < threadNum; i++ {
		go Scan(taskChan, wg)
	}

	// 生产者，不断地往taskChan channel发送数据，直接channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()
}

func Scan(taskChan chan map[string]int, wg *sync.WaitGroup) {
	// 每个协程都从channel中读取数据后开始扫描并入库
	for task := range taskChan {
		vars.ProgressBarPort.Increment()
		for ip, port := range task {
			slog.Debugf("Scanning %s:%d", ip, port)
			ip1, port1, err2 := Connect(ip, port)
			err := SaveResult(ip1, port1, err2)
			_ = err
			wg.Done()
		}
	}
	//vars.ProcessBarActive.Finish()
}

// Store 存储端口扫描结果到文件
func Store(ip string, ports []int) {
	filePort, err := os.Create("./result/portScan.txt")
	fileCrack, err := os.Create("./result/crackPorts.txt")
	if err != nil {
		slog.Error("Create file error", err)
	}

	for _, port := range ports {
		_, _ = filePort.WriteString(fmt.Sprintf("%v:%v\n", ip, port))
	}
	//存储可以被爆破的端口到crackPorts.txt文件中
	for _, port := range ports {
		protocol, ok := vars.PortNames[port]
		if ok && vars.SupportProtocols[protocol] {
			_, _ = fileCrack.WriteString(fmt.Sprintf("%v:%v\n", ip, port))
		}
	}
	vars.Result.Store(ip, ports)
}

func SaveResult(ip string, port int, err error) error {
	if err != nil {
		return err
	}
	v, ok := vars.Result.Load(ip)
	if ok {
		ports, ok1 := v.([]int)
		if ok1 {
			//判断结果是否已经保存
			flag := false
			for _, tmp := range ports {
				if tmp == port {
					flag = true
					break
				}
			}
			if !flag {
				ports = append(ports, port)
				Store(ip, ports)
			}
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		Store(ip, ports)
	}
	return err
}

func PrintResult() {
	vars.ProgressBarPort.Finish()
	vars.Result.Range(func(key, value interface{}) bool {
		fmt.Printf("ip:%v\n", key)
		fmt.Printf("ports: %v\n", value)
		fmt.Println(strings.Repeat("-", 100))
		return true
	})
}
