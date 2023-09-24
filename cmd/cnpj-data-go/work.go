package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/elizandrodantas/cnpj-data-go/database"
	"github.com/elizandrodantas/cnpj-data-go/external/download"
	"github.com/elizandrodantas/cnpj-data-go/external/gov"
	"github.com/elizandrodantas/cnpj-data-go/internal/filesystem"
	"github.com/elizandrodantas/cnpj-data-go/internal/migration"
	"github.com/elizandrodantas/cnpj-data-go/internal/unzip"

	"github.com/jmoiron/sqlx"
)

func startWork(conn *database.DatabaseConnectOptions, f *flags) {
	if f.files != "" {
		workWithPath(f.files, conn)
		return
	}

	workWithDownload(conn)
}

func workWithDownload(c *database.DatabaseConnectOptions) {
	Logger.Debug("starting working with download mode")

	resources, err := gov.Gov()
	if err != nil {
		Logger.Errorf("error while listing gov resources: %v\n", err)
		os.Exit(0)
	}

	Logger.Debugf(
		"gov resources found successfully | total of %s\n",
		ColorBoldInfo(len(resources.Details)))

	for _, resource := range resources.Details {
		if strings.HasSuffix(resource.Recurso.Link, ".zip") {
			println()

			Logger.Infof("Iniciando migração do recurso %s\n", resource.Name)

			timeStart := time.Now()

			table, ok := getTableNameWithFileName(resource.Name)
			if !ok {
				Logger.Error("could not get table name by resource name")
				continue
			}

			Logger.Debugf(
				"work on the resource %s and the %s table\n",
				ColorBoldInfo(resource.Name), ColorBoldInfo(table))

			client, err := database.Connect(c)
			if err != nil {
				Logger.Errorf("error create database client: %v\n", err)
				return
			}
			defer client.Close()

			Logger.Debugf(
				"starting download at %s\n",
				ColorBoldInfo(resource.Recurso.Link))

			processDownload, err := download.Start(resource.Recurso.Link)
			if err != nil {
				Logger.Errorf("error download %s\n", resource.Recurso.Link)
				return
			}
			defer filesystem.Delete(processDownload.GetName())

			Logger.Debugf(
				"performing migration of the %s resource in the %s table\n",
				ColorBoldInfo(resource.Name), ColorBoldInfo(table))

			Logger.Infof("Executando a migração do recurso %s na tabela %s\n", resource.Name, table)

			err = execute(client, table, processDownload.GetName())
			if err != nil {
				Logger.Errorf("error execute work: %v\n", err)
				return
			}

			Logger.Successf("Migração foi executada com sucesso na tabela %s\n", table)

			Logger.Debugf("migration was successful in the %s table\n", table)

			summary(timeStart, time.Now(), fmt.Sprintf("migração na tabela %s", table), nil)
		}

	}

}

func workWithPath(p string, c *database.DatabaseConnectOptions) {
	Logger.Debug("starting working with path mode")

	if !filesystem.ExistFile(p) {
		Logger.Error("dir path not found")
		os.Exit(0)
	}

	if !filesystem.IsDir(p) {
		Logger.Errorf("%s is not a directory\n", p)
		os.Exit(0)
	}

	dirList, err := filesystem.ReadDir(p)
	if err != nil {
		Logger.Errorf("error list dir: %v\n", err)
		os.Exit(0)
	}

	Logger.Debugf(
		"directory %s was listed and found %s files\n",
		ColorBoldInfo(p), ColorBoldInfo(len(dirList)))

	onlyZip := filesystem.OnlyZips(dirList)
	if len(onlyZip) == 0 {
		Logger.Errorf("no zip file found in directory %s\n", p)
		os.Exit(0)
	}

	Logger.Debugf(
		"in zip condition, %s files were found\n",
		ColorBoldInfo(len(onlyZip)))

	for _, k := range onlyZip {
		println()

		Logger.Infof("Iniciando migração no arquivo %s\n", k)

		timeStart := time.Now()

		client, err := database.Connect(c)
		if err != nil {
			Logger.Errorf("error create database client: %v\n", err)
			return
		}
		defer client.Close()

		zipPath := path.Join(p, k)
		table, ok := getTableNameWithFileName(k)
		if !ok {
			Logger.Errorf("table not found with filename %s\n", k)
			continue
		}

		Logger.Debugf(
			"performing migration of the %s resource in the %s table\n",
			ColorBoldInfo(k), ColorBoldInfo(table))

		Logger.Infof("Executando a migração do arquivo %s na tabela %s\n", k, table)

		err = execute(client, table, zipPath)
		if err != nil {
			Logger.Errorf("error execute work: %v\n", err)
			return
		}

		Logger.Successf("Migração foi executada com sucesso na tabela %s\n", table)

		Logger.Debugf("migration was successful in the %s table\n", table)

		summary(timeStart, time.Now(), fmt.Sprintf("migração na tabela %s", table), nil)
	}

}

func execute(client *sqlx.DB, table string, zipPath string) (err error) {
	if !filesystem.ExistFile(zipPath) {
		err = fmt.Errorf("zip path %s not found", zipPath)
		return
	}

	Logger.Debugf(
		"descompactor %s file\n",
		ColorBoldInfo(zipPath))

	unziped, err := unzip.NewUnzip(zipPath)
	if err != nil {
		return
	}
	defer filesystem.DeleteMany(unziped)

	Logger.Debugf(
		"starting the migration of data to %s table\n",
		ColorBoldInfo(table))

	task := migration.NewMigration(client, table, unziped[0])
	err = task.Execute()

	return
}
