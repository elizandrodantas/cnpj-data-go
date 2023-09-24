package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/elizandrodantas/cnpj-data-go/database"
	"github.com/elizandrodantas/cnpj-data-go/internal/box"
)

var tablesName = []string{
	"empresas",
	"estabelecimentos",
	"cnaes",
	"motivos",
	"municipios",
	"naturezas",
	"paises",
	"qualificacoes",
	"simples",
	"socios",
}

var driverName = []string{
	"sqlite3",
	"mysql",
	"postgres",
}

func getTableNameWithFileName(fileName string) (string, bool) {
	name := strings.ToLower(fileName)

	for _, k := range tablesName {
		if strings.Contains(name, k) {
			return k, true
		}
	}

	return "", false
}

func verifyExistTables(opt *database.DatabaseConnectOptions) error {
	client, err := database.Connect(opt)
	if err != nil {
		return err
	}
	defer client.Close()

	err = database.CreateTables(client)
	return err
}

func verifyDriver(f *flags) bool {
	drive := strings.ToLower(f.driverName)
	for _, k := range driverName {
		if strings.Contains(drive, k) {
			f.driverName = drive
			return true
		}
	}

	return false
}

func createConfigConnectionDatabase(f *flags) (c *database.DatabaseConnectOptions) {
	c = &database.DatabaseConnectOptions{
		DriverName:   f.driverName,
		DatabaseName: f.databaseName,
	}

	if f.driverName == "sqlite3" {
		return
	}

	if f.host == "" {
		f.host = "localhost"
	}

	c = &database.DatabaseConnectOptions{
		DriverName:   f.driverName,
		DatabaseName: f.databaseName,
		Host:         f.host,
		User:         f.user,
		Pass:         f.pass,
		Port:         f.port,
		SslMode:      f.sslMode,
	}

	return
}

func summary(timeStart, timeEnd time.Time, task string, a map[string]string) {
	box.New("Resumo")

	fmt.Printf("%s: %s\n", "Nome da tarefa", task)
	fmt.Printf("%s: %s\n", "Tempo total", timeEnd.Sub(timeStart))

	if len(a) > 0 {
		for key, value := range a {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	fmt.Print("\n")
}
