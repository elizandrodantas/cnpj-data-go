package tool

import (
	"fmt"
	"os"
)

type read struct {
	data []byte
}

func NewRead(path string) (*read, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return &read{}, err
	}

	if stat.IsDir() {
		return &read{}, fmt.Errorf("this path does not belong to a file")
	}

	out := readFile(path)
	if out == nil {
		return &read{}, fmt.Errorf("error in reading process")
	}

	return &read{out}, nil
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	return data
}

func (r *read) GetData() []byte {
	return r.data
}

func (r *read) GetDataString() string {
	return string(r.data)
}
