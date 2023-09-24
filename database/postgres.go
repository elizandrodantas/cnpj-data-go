package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func pgCreateClient(conf *DatabaseConnectOptions) (client *sqlx.DB, err error) {
	sslMode := "disable"

	if conf.SslMode {
		sslMode = "require"
	}

	client, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", conf.Host, conf.Port, conf.User, conf.DatabaseName, conf.Pass, sslMode))
	if err != nil {
		return
	}

	err = client.Ping()

	return
}
