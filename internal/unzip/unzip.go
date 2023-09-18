package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func NewUnzip(p string) ([]string, error) {
	output := []string{}

	if _, err := os.Stat(p); err != nil {
		return output, fmt.Errorf("zip path not found")
	}

	if !strings.HasSuffix(p, ".zip") {
		return output, fmt.Errorf("file is not zip type")
	}

	zipFile, err := zip.OpenReader(p)

	if err != nil {
		return output, err
	}

	defer zipFile.Close()

	for _, z := range zipFile.File {
		if z.Mode().IsDir() {
			continue
		}

		fileRead, err := z.Open()
		if err == nil {
			defer fileRead.Close()

			pathDest := path.Join(os.TempDir(), z.Name)

			newFile, err := os.Create(pathDest)
			if err != nil {
				return output, err
			}
			defer newFile.Close()

			_, err = io.Copy(newFile, fileRead)
			if err != nil {
				return output, err
			}

			output = append(output, pathDest)
		}
	}

	return output, nil
}
