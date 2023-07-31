package tool

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type files struct{}

func NewFiles() *files {
	return &files{}
}

type ReadFileStruct struct {
	data []byte
}

func (f *files) Read(path string) (*ReadFileStruct, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return &ReadFileStruct{}, err
	}

	if stat.IsDir() {
		return &ReadFileStruct{}, fmt.Errorf("this path does not belong to a file")
	}

	out := readFile(path)
	if out == nil {
		return &ReadFileStruct{}, fmt.Errorf("error in reading process")
	}

	return &ReadFileStruct{out}, nil
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	return data
}

func (r *ReadFileStruct) GetData() []byte {
	return r.data
}

func (r *ReadFileStruct) GetDataString() string {
	return string(r.data)
}

func (f *files) Delete(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	err := os.Remove(path)

	return err
}

func (f *files) DeleteMany(paths []string) error {
	for _, k := range paths {
		if _, err := os.Stat(k); err != nil {
			return err
		}

		err := os.Remove(k)

		if err != nil {
			return err
		}
	}

	return nil
}

func (f *files) ReadDir(directory string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return []fs.DirEntry{}, err
	}

	return files, nil
}

func (f *files) OnlyZips(in []fs.DirEntry) []string {
	output := []string{}

	for _, value := range in {
		if strings.HasSuffix(value.Name(), ".zip") {
			output = append(output, value.Name())
		}
	}

	return output
}
