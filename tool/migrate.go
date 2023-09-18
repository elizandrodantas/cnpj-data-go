package tool

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	BATCH_INSERT = 1500
)

type migrate struct {
	client *sqlx.DB
	table  string
	data   []byte
}

func NewMigrate(client *sqlx.DB, table string, data []byte) (*migrate, error) {
	if err := client.Ping(); err != nil {
		return &migrate{}, err
	}

	if data == nil {
		return &migrate{}, fmt.Errorf("data is required to migrate")
	}

	return &migrate{
		client: client,
		table:  table,
		data:   data,
	}, nil
}

func (m *migrate) prepare(sdata []string) []string {
	output := []string{}

	for _, v := range sdata {
		value := "("

		data := replaceData(v)

		value += data
		value += ")"

		output = append(output, value)
	}

	return output
}

func (m *migrate) Execute() (int64, error) {
	var (
		success = 0
		batch   = 0
	)

	sdpl := splitLineData(string(m.data))
	sdpl = removeNil(sdpl)

	for {
		last := batch + BATCH_INSERT

		if last > len(sdpl) {
			last = last - (last - len(sdpl))
		}

		preArr := m.prepare(sdpl[batch:last])
		preStr := strings.Join(preArr, ",")
		row := fmt.Sprintf("INSERT INTO %s VALUES %s", m.table, preStr)

		res, err := m.client.Exec(row)
		if err != nil {
			return -1, err
		}

		count, err := res.RowsAffected()
		if err != nil {
			return -1, err
		}

		updateProgressMigrate(int64(last), int64(len(sdpl)))

		success += int(count)
		batch += BATCH_INSERT

		if batch >= len(sdpl) {
			break
		}
	}

	fmt.Println()
	return int64(success), nil
}

func splitLineData(data string) []string {
	return strings.Split(data, "\n")
}

func replaceData(line string) string {
	return strings.ReplaceAll(line, ";", ",")
}

func removeNil(d []string) []string {
	var output []string

	for _, v := range d {
		if len(v) > 1 {
			byt := []byte{}

			for _, b := range []byte(v) {
				if b != 0 {
					byt = append(byt, b)
				}
			}

			output = append(output, string(byt))
		}
	}

	return output
}

func updateProgressMigrate(now, total int64) {
	percent := float64(now) / float64(total)
	progress := int(percent * 75.0)
	fmt.Printf("\rMigrate Progress: [")
	for i := 0; i < 75; i++ {
		if i < progress {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %.0f%%", percent*100)
}
