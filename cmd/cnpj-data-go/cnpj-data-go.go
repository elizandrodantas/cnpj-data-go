package main

import (
	"fmt"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/database"
	"github.com/elizandrodantas/cnpj-data-go/tool"
)

var (
	Logger = tool.Logger()
)

func Process(opt *database.DatabaseConnectOptions) {
	client, err := database.Connect(opt)

	if err != nil {
		Logger.Error(
			fmt.Sprintf("database connection error: %s", err.Error()),
		)

		return
	}
	defer client.Close()

	err = database.CreateTables(client)

	if err != nil {
		Logger.Error(
			fmt.Sprintf("error creating table: %s", err.Error()),
		)

		return
	}

	gov := tool.NewGov()

	Logger.Info("looking for resources in the gov")

	resources, err := gov.ListResources()

	if err != nil {
		Logger.Error(
			fmt.Sprintf("error to list resources of gov: %s", err.Error()),
		)

		return
	}

	Logger.Info("starting range in gov resources")

	for _, resource := range resources.Details {
		link := resource.Recurso.Link

		if strings.HasSuffix(link, ".zip") {
			Logger.Info(
				fmt.Sprintf("new zip link found | link: %s", link),
			)

			fileName := getFileName(link)
			download := tool.NewDownload()

			Logger.Info(
				fmt.Sprintf("starting download on file %s", fileName),
			)

			err := download.Start(link)
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error downloading %s | Message: %s", fileName, err.Error()),
				)

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}
			defer download.Delete(download.GetZipPath())

			Logger.Success(
				fmt.Sprintf("file %s was successfully downloaded", fileName),
			)

			Logger.Info("unzipping downloaded file")

			unziped, err := download.Unzip()
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error when unzipping file %s | Message: %s", fileName, err.Error()),
				)

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}
			defer download.DeleteMany(unziped)

			pathUnziped := unziped[0]

			Logger.Success(
				fmt.Sprintf("%d files were successfully unzipped", len(unziped)),
			)

			Logger.Info("starting reading")

			read, err := tool.NewRead(pathUnziped)
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error reading unzipped file | Message: %s", err.Error()),
				)

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			Logger.Success("the reading has been completed")

			tableName := getTableNameOfFileName(fileName)

			if tableName == "" {
				Logger.Error("non-gov files")

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			Logger.Info("starting the migration")

			migrate, err := tool.NewMigrate(client, tableName, read.GetData())
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error starting migration | Message: %s", err.Error()),
				)

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			Logger.Success("migration started successfully")

			Logger.Info(
				fmt.Sprintf("starting migration run on table %s", tableName),
			)

			Logger.Info("this could take a few minutes or even a few hours.")
			Logger.Info("wait...")

			lenRow, err := migrate.Execute()
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error when performing migration on table %s | Message: %s", tableName, err.Error()),
				)

				Logger.Info("\nprocess will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			Logger.Success(
				fmt.Sprintf("migration completed successfully on table %s", tableName),
			)
			Logger.Success(
				fmt.Sprintf("%d data were migrated %s", lenRow, tableName),
			)

			Logger.Info("this process has come to an end.")
		}
	}

	Logger.Success("finishing the processing")
}

func getFileName(link string) string {
	sb := strings.Split(link, "/")

	return sb[len(sb)-1]
}

func getTableNameOfFileName(fileName string) string {
	if strings.HasPrefix(fileName, "Empresas") {
		return "empresas"
	}

	if strings.HasPrefix(fileName, "Cnaes") {
		return "cnaes"
	}

	if strings.HasPrefix(fileName, "Estabelecimentos") {
		return "estabelecimentos"
	}

	if strings.HasPrefix(fileName, "Motivos") {
		return "motivos"
	}

	if strings.HasPrefix(fileName, "Municipios") {
		return "municipios"
	}

	if strings.HasPrefix(fileName, "Naturezas") {
		return "naturezas"
	}

	if strings.HasPrefix(fileName, "Paises") {
		return "paises"
	}

	if strings.HasPrefix(fileName, "Qualificacoes") {
		return "qualificacoes"
	}

	if strings.HasPrefix(fileName, "Simples") {
		return "simples"
	}

	if strings.HasPrefix(fileName, "Socios") {
		return "socios"
	}

	return ""
}
