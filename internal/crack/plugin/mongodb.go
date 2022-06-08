package plugin

import (
	"gopkg.in/mgo.v2"

	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
)

func ScanMongodb(s model.Service) (result model.ScanResult, err error) {
	result.Service = s
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", s.Username, s.Password, s.Ip, s.Port, "test")
	session, err := mgo.DialWithTimeout(url, vars.TimeOut)

	if err == nil {
		defer session.Close()
		err = session.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return result, err
}
