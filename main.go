package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/database"
	"github.com/elizandrodantas/cnpj-data-go/tool"
)

func main() {
	// client, err := database.Connect(&database.DatabaseConnectOptions{
	// 	DriverName: "sqlite3",
	// })

	// if err != nil {
	// 	log.Fatal("error connect database:", err.Error())
	// }

	// err = database.CreateTables(client)

	// read, _ := tool.NewRead("C:/Users/danta/AppData/Local/Temp/K3241.K03200Y0.D30708.EMPRECSV")
	// read, _ := tool.NewRead("C:/Users/danta/AppData/Local/Temp/F.K03200$Z.D30708.CNAECSV")
	// read, _ := tool.NewRead("C:/Users/danta/Desktop/K3241.K03200Y0.D30708.ESTABELE")
	// migrate, _ := tool.NewMigrate(client, "estabelecimentos", read.GetData())

	// result := migrate.Prepare(
	// 	[]string{
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 		`"abc";"def";"ghf"`,
	// 	},
	// )

	// _, err = migrate.Execute()

	// fmt.Println(err)
}

func getFileName(link string) string {
	sb := strings.Split(link, "/")

	return sb[len(sb)-1]
}

func getDatabaseNameOfFileName(fileName string) string {
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

func FirstInsert() {
	client, _ := database.Connect(&database.DatabaseConnectOptions{
		DriverName: "sqlite3",
	})

	database.CreateTables(client)

	read, err := tool.NewRead("C:/Users/danta/AppData/Local/Temp/F.K03200$Z.D30708.CNAECSV")
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := tool.NewMigrate(client, "cnaes", read.GetData())
	if err != nil {
		log.Fatal(err)
	}

	// migrate.Prepare()

	rows, err := migrate.Execute()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Foram inserdos", rows)
}

func TestRead() {
	read, err := tool.NewRead("C:/Users/danta/AppData/Local/Temp/F.K03200$Z.D30708.CNAECSV")
	if err != nil {
		log.Fatal(err)
	}

	str := read.GetDataString()

	fmt.Println(str)
}

func TestListGov() {
	gov := tool.NewGov()

	list, err := gov.ListResources()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}

func DownloadTest() {
	fmt.Println()
	dow := tool.NewDownload()

	err := dow.Start("https://dadosabertos.rfb.gov.br/CNPJ/Estabelecimentos0.zip")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("baixado com sucesso")

	res, err := dow.Unzip()

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("zipped success", res)
}

func ProcessAll() {
	client, err := database.Connect(&database.DatabaseConnectOptions{
		DriverName: "sqlite3",
	})

	if err != nil {
		log.Fatal("error connect database:", err.Error())
	}

	err = database.CreateTables(client)

	if err != nil {
		log.Fatalln("error to create tables:", err.Error())
	}

	gov := tool.NewGov()

	resources, err := gov.ListResources()

	if err != nil {
		log.Fatalln("error to list resources of gov:", err.Error())
	}

	log.Println("starting range in resources")

	for _, resource := range resources.Details {
		link := resource.Recurso.Link

		if strings.HasSuffix(link, ".zip") {
			log.Println("new zip file found, initiating database download, [", link, "]")

			fileName := getFileName(link)

			log.Println("starting download of", fileName)
			download := tool.NewDownload()

			err := download.Start(link)

			if err != nil {
				log.Println("error downloading", fileName, "| Error:", err.Error())
				log.Println("process will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			unziped, err := download.Unzip()
			if err != nil {
				log.Println("error when unzipping", fileName, "| Error:", err.Error())
				log.Println("process will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			fmt.Println(unziped)

			pathUnizped := unziped[0]
			fmt.Println(pathUnizped)
			read, err := tool.NewRead(pathUnizped)
			if err != nil {
				log.Println("error reading file", fileName, "| Error:", err.Error())
				log.Println("process will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			databaseName := getDatabaseNameOfFileName(fileName)
			migrate, err := tool.NewMigrate(client, databaseName, read.GetData())
			if err != nil {
				log.Println("error migrate data to", "", "| Error:", err.Error())
				log.Println("process will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			// migrate.Prepare()
			lenRows, err := migrate.Execute()
			if err != nil {
				log.Println("error migrate data to", databaseName, "| Error:", err.Error())
				log.Println("process will be terminated and overthrown")
				database.DropTables(client)
				break
			}

			log.Println("migrated successful to", databaseName, "| Total migrated:", lenRows)
		}
	}
}
