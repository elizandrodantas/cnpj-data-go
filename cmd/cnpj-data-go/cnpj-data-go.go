package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/database"
	"github.com/elizandrodantas/cnpj-data-go/tool"
)

var (
	Logger = tool.Logger()
)

func DropTables(opt *database.DatabaseConnectOptions) {
	Logger.Info("starting drop tables")

	client, err := database.Connect(opt)
	if err != nil {
		Logger.Error(
			fmt.Sprintf("error connecting to database: %s", err.Error()),
		)
		return
	}
	defer client.Close()

	var confirmation string

	fmt.Print("are you sure you want to delete the tables?? (yes/no) ")
	fmt.Scan(&confirmation)

	if strings.ToLower(confirmation) != "yes" {
		Logger.Info("deletion of tables canceled")
		return
	}

	err = database.DropTables(client)

	if err != nil {
		Logger.Error(
			fmt.Sprintf("error deleting tables: %s", err.Error()),
		)
		return
	}

	Logger.Success("successfully deleted tables")
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

func DownloadProcess(opt *database.DatabaseConnectOptions) {
	Logger.Info("starting process with download")

	gov := tool.NewGov()

	Logger.Info("starting search for data in gov")

	resources, err := gov.ListResources()
	if err != nil {
		Logger.Error(
			fmt.Sprintf("error to list resources of gov: %s", err.Error()),
		)
		return
	}
	Logger.Success("gov data successfully retrieved")

	Logger.Info("starting download and migration of resources")

	err = verifyExistTables(opt)

	if err != nil {
		Logger.Error(
			fmt.Sprintf("error checking tables: %s", err.Error()),
		)
		return
	}

	for _, resource := range resources.Details {
		link := resource.Recurso.Link

		if strings.HasSuffix(link, ".zip") {
			Logger.Info(
				fmt.Sprintf("new zip resource found | id: %s | created: %s", resource.ID, resource.Created),
			)

			client, err := database.Connect(opt)
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error connecting to database: %s", err.Error()),
				)
				break
			}
			defer client.Close()
			Logger.Success("successfully connected to the database")

			Logger.Info(
				fmt.Sprintf("starting download > %s", link),
			)

			download := tool.NewDownload()
			err = download.Start(link)
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error downloading resource: %s", err.Error()),
				)
				continue
			}
			defer tool.NewFiles().Delete(download.GetZipPath())
			Logger.Success(
				fmt.Sprintf("resource downloaded successfully | Name: %s | ID: %s", resource.Name, resource.ID),
			)

			Logger.Info(
				fmt.Sprintf("unzipping resource | ID: %s", resource.ID),
			)

			unziped, err := tool.NewUnzip(download.GetZipPath())
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error unzipping resource: %s", err.Error()),
				)
				break
			}
			defer tool.NewFiles().DeleteMany(unziped)
			Logger.Success(
				fmt.Sprintf("resource was successfully unzipped | ID: %s", resource.ID),
			)

			Logger.Info(
				fmt.Sprintf("starting to read the resource | ID: %s", resource.ID),
			)
			read, err := tool.NewFiles().Read(unziped[0])
			if err != nil {
				Logger.Error(
					fmt.Sprintf("database read error: %s", err.Error()),
				)
				break
			}
			Logger.Success(
				fmt.Sprintf("resource reading completed successfully | ID: %s", resource.ID),
			)

			table := getTableNameOfFileName(resource.Name)

			if table == "" {
				Logger.Error("no bank found for this resource")
				continue
			}

			Logger.Info(
				fmt.Sprintf("starting feature migration | ID: %s", resource.ID),
			)

			migrate, err := tool.NewMigrate(client, table, read.GetData())
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error instantiating migration: %s", err.Error()),
				)
				continue
			}

			total, err := migrate.Execute()
			if err != nil {
				Logger.Error(
					fmt.Sprintf("error when migrating: %s", err.Error()),
				)
				continue
			}
			Logger.Success(
				fmt.Sprintf("full migration on table %s | Total: %d | Resource ID: %s | Resource Name: %s", table, total, resource.ID, resource.Name),
			)

			fmt.Println()
		}
	}
}

func FilesProcess(directory string, opt *database.DatabaseConnectOptions) {
	Logger.Info("starting process with files")

	if stat, err := os.Stat(directory); err != nil || !stat.IsDir() {
		Logger.Error(
			fmt.Sprintf("directory '%s' not found", directory),
		)
		return
	}

	dirList, err := tool.NewFiles().ReadDir(directory)
	if err != nil {
		Logger.Error(
			fmt.Sprintf("error listing directory: %s", err.Error()),
		)
	}

	onlyZip := tool.NewFiles().OnlyZips(dirList)

	if len(onlyZip) == 0 {
		Logger.Info(
			fmt.Sprintf("no resource was found in the directory (%s)", directory),
		)
		return
	}

	Logger.Info(
		fmt.Sprintf("-> %d resources were found in the assets directories %s", len(onlyZip), directory),
	)

	err = verifyExistTables(opt)

	if err != nil {
		Logger.Error(
			fmt.Sprintf("error checking tables: %s", err.Error()),
		)
		return
	}

	Logger.Info("starting the migration")

	for _, zip := range onlyZip {
		client, err := database.Connect(opt)
		if err != nil {
			Logger.Error(
				fmt.Sprintf("error connecting to database: %s", err.Error()),
			)
			break
		}
		defer client.Close()
		Logger.Success("successfully connected to the database")

		zipPath := path.Join(directory, zip)
		table := getTableNameOfFileName(zip)

		if table == "" {
			Logger.Info(
				fmt.Sprintf("no table found for zip %s", zip),
			)
			continue
		}

		Logger.Info(
			fmt.Sprintf("unzipping resource | Zip: %s", zip),
		)

		unziped, err := tool.NewUnzip(zipPath)
		if err != nil {
			Logger.Error(
				fmt.Sprintf("error unzipping resource: %s", err.Error()),
			)
			break
		}
		defer tool.NewFiles().DeleteMany(unziped)

		Logger.Success(
			fmt.Sprintf("resource was successfully unzipped | Zip: %s", zip),
		)

		Logger.Info(
			fmt.Sprintf("starting to read the resource | Zip: %s", zip),
		)
		read, err := tool.NewFiles().Read(unziped[0])
		if err != nil {
			Logger.Error(
				fmt.Sprintf("database read error: %s", err.Error()),
			)
			break
		}
		Logger.Success(
			fmt.Sprintf("resource reading completed successfully | Zip: %s", zip),
		)

		Logger.Info(
			fmt.Sprintf("starting feature migration | Zip: %s", zip),
		)

		migrate, err := tool.NewMigrate(client, table, read.GetData())
		if err != nil {
			Logger.Error(
				fmt.Sprintf("error instantiating migration: %s", err.Error()),
			)
			continue
		}

		total, err := migrate.Execute()
		if err != nil {
			Logger.Error(
				fmt.Sprintf("error when migrating: %s", err.Error()),
			)
			continue
		}
		Logger.Success(
			fmt.Sprintf("full migration on table %s | Total: %d | Resource Zip: %s", table, total, zip),
		)

		fmt.Println()
	}
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
