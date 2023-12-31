package filesystem

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

func ReadDir(directory string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return []fs.DirEntry{}, err
	}

	return files, nil
}

func OnlyZips(in []fs.DirEntry) []string {
	output := []string{}

	for _, value := range in {
		if strings.HasSuffix(value.Name(), ".zip") {
			output = append(output, value.Name())
		}
	}

	return output
}

func LenLines(name string) int64 {
	file, err := OpenFile(name)
	if err != nil {
		return -1
	}

	len := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		len++
	}

	if err := scanner.Err(); err != nil {
		return -1
	}

	return int64(len)
}

func Delete(name string) error {
	if _, err := Stat(name); err != nil {
		return err
	}

	err := os.Remove(name)
	return err
}

func DeleteMany(paths []string) error {
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

func OpenFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return file, err
}

func Stat(name string) (fs.FileInfo, error) {
	file, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	return file, err
}

func ExistFile(name string) bool {
	if _, err := Stat(name); err != nil {
		return false
	}

	return true
}

func IsDir(name string) bool {
	file, err := Stat(name)
	if err != nil {
		return false
	}

	return file.IsDir()
}
