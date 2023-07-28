package database

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConnectOptions struct {
	DriverName   string
	Host         string
	User         string
	Pass         string
	DatabaseName string
	Port         int
	SslMode      bool
}

func Connect(opt *DatabaseConnectOptions) (*sqlx.DB, error) {
	sslMode := "disable"

	if opt.DatabaseName == "" {
		opt.DatabaseName = "database"
	}

	if opt.SslMode {
		sslMode = "require"
	}

	dataSource := ""

	if strings.HasPrefix(opt.DriverName, "sqlite3") {
		dataSource = fmt.Sprintf("file:%s.sqlite", opt.DatabaseName)
	} else {
		dataSource = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", opt.Host, opt.Port, opt.User, opt.DatabaseName, opt.Pass, sslMode)
	}

	client, err := sqlx.Connect(opt.DriverName, dataSource)

	return client, err
}

func CreateTables(client *sqlx.DB) (err error) {
	err = client.Ping()

	if err != nil {
		return
	}

	_, err = client.Exec(CREATE_TABLES)

	return
}

func DropTables(client *sqlx.DB) (err error) {
	err = client.Ping()

	if err != nil {
		return
	}

	_, err = client.Exec(DROP_TABLES)

	return
}
