package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func mysqlCreateClient(conf *DatabaseConnectOptions) (client *sqlx.DB, err error) {
	client, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.User, conf.Pass, conf.Host, conf.DatabaseName))
	if err != nil {
		return
	}

	err = client.Ping()

	return
}
