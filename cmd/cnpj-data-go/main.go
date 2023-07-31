package main

import (
	"flag"

	"github.com/elizandrodantas/cnpj-data-go/database"
)

type flags struct {
	driverName   string
	host         string
	user         string
	pass         string
	databaseName string
	port         int
	sslMode      bool
	files        string
	drop         bool
}

func main() {
	var f flags

	flag.StringVar(&f.driverName, "D", "sqlite3", "driver name (sqlite3 or mysql or postgres)")
	flag.StringVar(&f.host, "h", "localhost", "database host (if not sqlite)")
	flag.StringVar(&f.user, "u", "", "database username (if not sqlite)")
	flag.StringVar(&f.pass, "p", "", "database password (if not sqlite)")
	flag.StringVar(&f.databaseName, "dbname", "database", "database name")
	flag.IntVar(&f.port, "P", 0, "database port (if not sqlite)")
	flag.BoolVar(&f.sslMode, "ssl", false, "active sslmode connection")
	flag.StringVar(&f.files, "F", "", "use files to migrate")
	flag.BoolVar(&f.drop, "drop", false, "delete tables from database")
	flag.Parse()

	if f.driverName != "sqlite3" && f.driverName != "mysql" && f.driverName != "postgres" {
		Logger.Error("drivername invalid, accepted drivenames are: sqlite3, mysql or postgres")
		return
	}

	var conn database.DatabaseConnectOptions

	if f.driverName == "sqlite3" {
		conn = database.DatabaseConnectOptions{
			DriverName:   f.driverName,
			DatabaseName: f.databaseName,
		}
	} else {
		if f.host == "" {
			f.host = "localhost"
		}

		conn = database.DatabaseConnectOptions{
			DriverName:   f.driverName,
			DatabaseName: f.databaseName,
			Host:         f.host,
			User:         f.user,
			Pass:         f.pass,
			Port:         f.port,
			SslMode:      f.sslMode,
		}
	}

	if f.drop {
		DropTables(&conn)
		return
	}

	if f.files != "" {
		FilesProcess(f.files, &conn)
	} else {
		DownloadProcess(&conn)
	}
}
