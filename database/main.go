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
	if strings.Contains(opt.DriverName, "postgres") {
		return pgCreateClient(opt)
	}

	if strings.Contains(opt.DriverName, "mysql") {
		return mysqlCreateClient(opt)
	}

	return sqliteCreateClient(opt)
}

func CreateTables(client *sqlx.DB) (err error) {
	err = client.Ping()
	if err != nil {
		return
	}

	for i := 0; i < len(CREATE_TABLES); i++ {
		_, err = client.Exec(CREATE_TABLES[i])

		if err != nil {
			return
		}
	}

	return
}

func DropTables(client *sqlx.DB) (err error) {
	err = client.Ping()

	if err != nil {
		return
	}

	for i := 0; i < len(TABLES_NAME_LIST); i++ {
		row := fmt.Sprintf(DROP_TABLES, TABLES_NAME_LIST[i])

		_, err = client.Exec(row)

		if err != nil {
			return
		}
	}

	return
}
