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
	} else if strings.HasPrefix(opt.DriverName, "mysql") {
		dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s", opt.User, opt.Pass, opt.Host, opt.DatabaseName)
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

	driver := client.DriverName()

	if driver == "mysql" {
		defer AlterEngine(client)
		defer AlterCharset(client)
	}

	for i := 0; i < len(CREATE_TABLES); i++ {
		_, err = client.Exec(CREATE_TABLES[i])

		if err != nil {
			return
		}
	}

	return
}

func AlterEngine(client *sqlx.DB) (err error) {
	err = client.Ping()

	if err != nil {
		return
	}

	for i := 0; i < len(TABLES_NAME_LIST); i++ {
		row := fmt.Sprintf(ALTER_ENGINE, TABLES_NAME_LIST[i], "InnoDB")
		_, err = client.Exec(row)

		if err != nil {
			return
		}
	}

	return nil
}

func AlterCharset(client *sqlx.DB) (err error) {
	err = client.Ping()

	if err != nil {
		return
	}

	for i := 0; i < len(TABLES_NAME_LIST); i++ {
		row := fmt.Sprintf(ALTER_CHARSET, TABLES_NAME_LIST[i], "InnoDB")
		_, err = client.Exec(row)

		if err != nil {
			return
		}
	}

	return nil
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
