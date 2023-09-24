package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func sqliteCreateClient(conf *DatabaseConnectOptions) (client *sqlx.DB, err error) {
	if conf.DatabaseName == "" {
		conf.DatabaseName = "database"
	}

	client, err = sqlx.Connect("sqlite3", fmt.Sprintf("file:%s.db", conf.DatabaseName))
	if err != nil {
		return
	}

	err = client.Ping()

	return
}
