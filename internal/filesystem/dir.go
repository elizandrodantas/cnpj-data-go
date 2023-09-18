package filesystem

import (
	"io/fs"
	"os"
	"strings"
)

func (f *file) ReadDir(directory string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return []fs.DirEntry{}, err
	}

	return files, nil
}

func (f *file) OnlyZips(in []fs.DirEntry) []string {
	output := []string{}

	for _, value := range in {
		if strings.HasSuffix(value.Name(), ".zip") {
			output = append(output, value.Name())
		}
	}

	return output
}
