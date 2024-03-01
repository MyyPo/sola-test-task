package defaultval

import "sola-test-task/pkg/log"

const (
	ServerHost = "0.0.0.0"
	ServerPort = "12499"
	LoggerMode = log.Debug
	LoggerPath = "/var/log/sola/app/server.log"

	DbHost         = "sola-database"
	DbPort         = "5432"
	DbSslMode      = "disable"
	DbName         = "sola"
	DbUserName     = "postgres"
	DbUserPassword = "postgres-sola"
)
