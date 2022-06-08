package plugin

import (
	"github.com/jlaffaye/ftp"

	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model"
	"github.com/EwanSunn/secScan/internal/pkg/model/vars"
)

func ScanFtp(s model.Service) (result model.ScanResult, err error) {
	result.Service = s
	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", s.Ip, s.Port), vars.TimeOut)
	if err == nil {
		err = conn.Login(s.Username, s.Password)
		if err == nil {
			defer func() {
				err = conn.Logout()
			}()
			result.Result = true
		}
	}
	return result, err
}

