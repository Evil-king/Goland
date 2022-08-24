package pdb

import (
	"time"

	"github.com/gogf/gf/database/gdb"
)

// DBConfig is the package db connection config
type DBConfig struct {
	Host             string        // [Must] Host of DB. Use ip or domain: 127.0.0.1, localhost.
	Port             string        // [Must] Port of DB.
	User             string        // [Must] Authentication username.
	Pass             string        // [Must] Authentication password.
	Database         string        // [Must] Database name.
	Type             string        // [Optional] DB type. Default to mysql. [mysql|sqlite|mssql|pgsql|oracle]
	MaxIdleConnCount int           // [Optional] Max idle connection configuration. Default to 100.
	MaxOpenConnCount int           // [Optional] Max open connection configuration. Default to 1100.
	MaxConnLifetime  time.Duration // [Optional] Max connection TTL configuration. Default to 3600s.
	Debug            bool          // (Optional) Debug mode enables debug information logging and output.
	Charset          string
}

// SetDB sets the global configuration for package and test the connection.
// Returns an error if connection is invalid.
func SetDB(config *DBConfig) error {
	// use default if unset
	if len(config.Type) == 0 {
		config.Type = DB_TYPE_MYSQL
	}
	if config.MaxIdleConnCount == 0 {
		config.MaxIdleConnCount = MAX_IDLE_CONN
	}
	if config.MaxOpenConnCount == 0 {
		config.MaxOpenConnCount = MAX_OPEN_CONN
	}
	if config.MaxConnLifetime == 0 {
		config.MaxConnLifetime = MAX_CONN_LIFETIME
	}
	if config.Charset == "" {
		config.Charset = "utf8mb4"
	}

	node := gdb.ConfigNode{
		Host:             config.Host,
		Port:             config.Port,
		User:             config.User,
		Pass:             config.Pass,
		Name:             config.Database,
		Type:             config.Type,
		MaxIdleConnCount: config.MaxIdleConnCount,
		MaxOpenConnCount: config.MaxOpenConnCount,
		Debug:            config.Debug,
		CreatedAt:        CREATED_TIME,
		UpdatedAt:        UPDATED_TIME,
		Charset:          config.Charset,
	}

	// set config
	gdb.SetConfig(gdb.Config{
		gdb.DefaultGroupName: gdb.ConfigGroup{
			node,
		},
	})

	// test connection
	_, err := gdb.Instance(gdb.DefaultGroupName)
	return err
}
