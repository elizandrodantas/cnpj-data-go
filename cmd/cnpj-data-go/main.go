package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/external/metrics"
	"github.com/elizandrodantas/cnpj-data-go/internal/logger"
	"github.com/fatih/color"
)

var (
	Logger        = logger.Logger{}
	ColorBoldInfo = color.New(color.Bold, color.FgYellow).SprintFunc()
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

	metric             bool
	verbose            bool
	ignoreCreateTables bool
}

func main() {
	var f flags

	flag.StringVar(&f.driverName, "D", "sqlite3", "driver name (sqlite3 or mysql or postgres)")

	flag.StringVar(&f.host, "host", "localhost", "database host (if not sqlite)")
	flag.StringVar(&f.user, "user", "", "database username (if not sqlite)")
	flag.StringVar(&f.pass, "pass", "", "database password (if not sqlite)")
	flag.StringVar(&f.databaseName, "dbname", "database", "database name")
	flag.IntVar(&f.port, "port", 0, "database port (if not sqlite)")
	flag.BoolVar(&f.sslMode, "ssl", false, "active sslmode connection")

	flag.StringVar(&f.files, "F", "", "use files to migrate")
	flag.BoolVar(&f.drop, "drop", false, "delete tables from database")

	flag.BoolVar(&f.metric, "metric", false, "open metrics runtime")
	flag.BoolVar(&f.verbose, "verbose", false, "shows detailed process information")
	flag.BoolVar(&f.ignoreCreateTables, "ignore-create-table", false, "ignores the creation of tables")

	flag.Parse()

	if f.verbose {
		Logger.DebugLogger = true
	}

	ok := verifyDriver(&f)
	if !ok {
		Logger.Errorf("drivername invalid, accepted drivenames are: %s\n", strings.Join(driverName, ", "))
		os.Exit(0)
	}

	if f.metric {
		metrics.Start()
	}

	c := createConfigConnectionDatabase(&f)

	if f.drop {
		dropTables(c)
		os.Exit(0)
	}

	if !f.ignoreCreateTables {
		err := verifyExistTables(c)
		if err != nil {
			Logger.Errorf("error verify tables condition: %v\n", err)
			os.Exit(0)
		}
	}

	runtime.GOMAXPROCS(35)

	startWork(c, &f)
}
