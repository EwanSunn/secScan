package model

import (
	"encoding/gob"
	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
	"github.com/EwanSunn/secScan/internal/pkg/slog"
	"github.com/EwanSunn/secScan/internal/pkg/util/hash"
	"github.com/patrickmn/go-cache"
	"os"
	"strings"
	"time"
)

func init() {
	gob.Register(Service{})
	gob.Register(ScanResult{})
}

func SaveResult(result ScanResult, err error) {
	if err == nil && result.Result {
		var k string
		protocol := strings.ToUpper(result.Service.Protocol)

		if protocol == "REDIS" {
			k = fmt.Sprintf("%v-%v-%v", result.Service.Ip, result.Service.Port, result.Service.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", result.Service.Ip, result.Service.Port, result.Service.Username)
		}

		h := hash.MakeTaskHash(k)
		hash.SetTaskHash(h)

		_, found := vars.CacheService.Get(k)
		if !found {
			slog.Infof("Ip: %v, Port: %v, Protocol: [%v], Username: %v, Password: %v", result.Service.Ip,
				result.Service.Port, result.Service.Protocol, result.Service.Username, result.Service.Password)
		}
		vars.CacheService.Set(k, result, cache.NoExpiration)
	}
}

func CacheStatus() (count int, items map[string]cache.Item) {
	count = vars.CacheService.ItemCount()
	items = vars.CacheService.Items()
	return count, items
}

func ResultTotal() {
	vars.ProgressBarPassword.Finish()
	slog.Info(fmt.Sprintf("Finshed scan, total result: %v, used time: %v",
		vars.CacheService.ItemCount(),
		time.Since(vars.StartTime)))
}

func SaveResultToFile() error {
	return vars.CacheService.SaveFile("./result/password_crack.db")
}

func DumpToFile(filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, items := CacheStatus()
	for _, v := range items {
		result := v.Object.(ScanResult)
		_, _ = file.WriteString(fmt.Sprintf("%v:%v|%v,%v:%v\n",
			result.Service.Ip,
			result.Service.Port,
			result.Service.Protocol,
			result.Service.Username,
			result.Service.Password),
		)
	}

	return err
}
