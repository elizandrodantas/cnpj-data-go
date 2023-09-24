package main

import (
	"os"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/database"
)

func dropTables(opt *database.DatabaseConnectOptions) {
	Logger.Debug("starting tables drop")

	client, err := database.Connect(opt)
	if err != nil {
		Logger.Errorf("error connecting to database: %v\n", err)
		os.Exit(0)
	}
	defer client.Close()

	confirmationExampleInput := "yes/no"
	confirmation := Logger.InputString("shows detailed process information", &confirmationExampleInput)

	if strings.Index(confirmation, "y") != 0 {
		Logger.Info("deletion of tables canceled")
		return
	}

	err = database.DropTables(client)
	if err != nil {
		Logger.Errorf("error deleting tables: %v\n", err)
		os.Exit(0)
	}

	Logger.Success("successfully deleted tables")
}
