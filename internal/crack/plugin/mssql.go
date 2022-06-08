package plugin

import (
	"database/sql"
	"fmt"
	"github.com/EwanSunn/secScan/internal/pkg/model"

	_ "github.com/denisenkom/go-mssqldb"
)

func ScanMssql(service model.Service) (result model.ScanResult, err error) {
	result.Service = service

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", service.Ip,
		service.Port, service.Username, service.Password, "master")

	db, err := sql.Open("mssql", dataSourceName)
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
