package util

import (
	"bufio"
	"github.com/EwanSunn/secScan/internal/config"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"net"
	"os"
	"strconv"
	"strings"
)

func ReadIpList(fileName string) (ipLists []net.IP, err error) {
	ipListFile, err := os.Open(fileName)
	if err != nil {
		config.Config.Log.Error("Open ip List file err, %v", err)
	}

	defer func() {
		if ipListFile != nil {
			_ = ipListFile.Close()
		}
	}()
	scanner := bufio.NewScanner(ipListFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ipList, err := GetIpList(line)
		if err != nil {
			config.Config.Log.Error("ReadIpList Error, %v", err)
		}
		for _, ip := range ipList {
			ipLists = append(ipLists, ip)
		}
	}
	return ipLists, err
}

func ReadIpPortList(fileName string) (ipList []model.IpAddr) {
	ipListFile, err := os.Open(fileName)
	if err != nil {
		config.Config.Log.Error("Open ipPort List file err, %v", err)
	}

	defer func() {
		if ipListFile != nil {
			_ = ipListFile.Close()
		}
	}()

	scanner := bufio.NewScanner(ipListFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ipPort := strings.TrimSpace(line)
		t := strings.Split(ipPort, ":")
		ip := t[0]
		portProtocol := t[1]
		tmpPort := strings.Split(portProtocol, "|")
		// ip列表中指定了端口对应的服务
		if len(tmpPort) == 2 {
			port, _ := strconv.Atoi(tmpPort[0])
			protocol := strings.ToUpper(tmpPort[1])
			if vars.SupportProtocols[protocol] {
				addr := model.IpAddr{Ip: ip, Port: port, Protocol: protocol}
				ipList = append(ipList, addr)
			} else {
				config.Config.Log.Infof("Not support %v, ignore: %v:%v", protocol, ip, port)
			}
		} else {
			// 通过端口查服务
			port, err := strconv.Atoi(tmpPort[0])
			if err == nil {
				protocol, ok := vars.PortNames[port]
				if ok && vars.SupportProtocols[protocol] {
					addr := model.IpAddr{Ip: ip, Port: port, Protocol: protocol}
					ipList = append(ipList, addr)
				}
			}
		}

	}

	return ipList
}

func ReadUserDict(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
		config.Config.Log.Fatalf("Open user dict file err, %v", err)
	}

	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users, err
}

func ReadPasswordDict(passDict string) (password []string, err error) {
	file, err := os.Open(passDict)
	if err != nil {
		config.Config.Log.Fatalf("Open password dict file err, %v", err)
	}

	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	password = append(password, "")
	return password, err
}
