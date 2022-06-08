package plugin

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/EwanSunn/secScan/internal/pkg/model"
)

func ScanPostgres(service model.Service) (result model.ScanResult, err error) {
	result.Service = service

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", service.Username,
		service.Password, service.Ip, service.Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)

	if err == nil {
		defer func() {
			err = db.Close()
		}()
		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return result, err
}