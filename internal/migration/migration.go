package migration

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/internal/filesystem"
	"github.com/elizandrodantas/cnpj-data-go/internal/progress"
	"github.com/jmoiron/sqlx"
)

var (
	BATCH_INSERT = 1500
)

type migration struct {
	clientDB *sqlx.DB

	file   string
	table  string
	length int64
}

func NewMigration(clientDB *sqlx.DB, table string, file string) *migration {
	lenFile := filesystem.LenLines(file)

	return &migration{clientDB, file, table, lenFile}
}

func migrate(client *sqlx.DB, data []string, table string) error {
	outNil := parseNil(data)
	parsed := parseInsertValues(outNil)

	row := fmt.Sprintf("INSERT INTO %s VALUES %s", table, strings.Join(parsed, ","))

	_, err := client.Exec(row)
	if err != nil {
		return err
	}

	return nil
}

func (m *migration) Execute() error {
	file, err := filesystem.OpenFile(m.file)
	if err != nil {
		return err
	}
	defer file.Close()

	err = m.clientDB.Ping()
	if err != nil {
		return fmt.Errorf("error ping client: %v", err)
	}

	dp := progress.NewProgressDefault("migrating", m.length)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var (
		batch = BATCH_INSERT
		index = 0

		data []string
	)

	for scanner.Scan() {
		data = append(data, scanner.Text())
		index++

		if index >= batch {
			err := migrate(m.clientDB, data, m.table)
			if err != nil {
				return err
			}

			dp.Add(index)
			index = 0
			data = nil
		}
	}

	if len(data) > 0 {
		err := migrate(m.clientDB, data, m.table)
		if err != nil {
			return err
		}

		dp.Add(len(data))
		index = 0
		data = nil
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	dp.Finish()

	return nil
}
