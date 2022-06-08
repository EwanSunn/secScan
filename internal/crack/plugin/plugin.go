package plugin

import "github.com/EwanSunn/secScan/internal/pkg/model"

type ScanFunc func(service model.Service) (result model.ScanResult, err error)

var (
	ScanFuncMap map[string]ScanFunc
)

func init() {
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["FTP"] = ScanFtp
	ScanFuncMap["SSH"] = ScanSsh
	ScanFuncMap["MYSQL"] = ScanMysql
	ScanFuncMap["MSSQL"] = ScanMssql
	ScanFuncMap["REDIS"] = ScanRedis
	ScanFuncMap["POSTGRESQL"] = ScanPostgres
	ScanFuncMap["MONGODB"] = ScanMongodb
}
