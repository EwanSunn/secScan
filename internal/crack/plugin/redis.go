package plugin

import (
	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"

	"github.com/go-redis/redis"
)

func ScanRedis(s model.Service) (result model.ScanResult, err error) {
	result.Service = s
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", s.Ip, s.Port),
		Password: s.Password, DB: 0, DialTimeout: vars.TimeOut}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	if err != nil {
		return result, err
	}

	result.Result = true

	defer func() {
		if client != nil {
			_ = client.Close()
		}
	}()

	return result, err
}

