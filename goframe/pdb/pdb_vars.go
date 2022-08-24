package pdb

import "time"

const (
	// const BD type
	DB_TYPE_MYSQL   = "mysql"
	DB_TYPE_SQLLITE = "sqlite"
	DB_TYPE_MSSQL   = "mssql"
	DB_TYPE_PGSQL   = "pgsql"
	DB_TYPE_ORACLE  = "oracle"

	// default configuration
	MAX_IDLE_CONN     = 100
	MAX_OPEN_CONN     = 1100
	MAX_CONN_LIFETIME = time.Second * time.Duration(3600)

	// auto time field rename
	CREATED_TIME = "created_time"
	UPDATED_TIME = "updated_time"
)
