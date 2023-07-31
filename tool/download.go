package tool

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Download struct {
	fileZipPath string
	downloaded  bool
	path        string
}

func NewDownload() *Download {
	return &Download{
		path: os.TempDir(),
	}
}

func (d *Download) Start(link string) error {
	if !strings.HasPrefix(link, "http") {
		return fmt.Errorf("this link is invalid")
	}

	if _, err := url.Parse(link); err != nil {
		return fmt.Errorf("link parse error")
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	response, err := client.Get(link)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	tempFile, err := os.CreateTemp("", "gov-br-*.zip")

	if err != nil {
		return err
	}
	defer tempFile.Close()

	d.fileZipPath = tempFile.Name()
	contentLength := response.ContentLength

	err = copyFile(response.Body, tempFile, contentLength)
	if err != nil {
		return err
	}

	d.downloaded = true

	return nil
}

func copyFile(body io.Reader, output *os.File, contentLength int64) error {
	buffer := make([]byte, 1024)
	totalByteRead := int64(0)

	for {
		n, err := body.Read(buffer)
		if n > 0 {
			totalByteRead += int64(n)

			updateProgressDownload(totalByteRead, contentLength)

			output.Write(buffer[:n])
		}

		if totalByteRead >= contentLength {
			break
		}

		if err == io.EOF {
			continue
		}

		if err != nil {
			return err
		}
	}

	fmt.Println()
	return nil
}

func updateProgressDownload(read, total int64) {
	percent := float64(read) / float64(total)
	progress := int(percent * 75.0)
	fmt.Printf("\rDownload Progress: [")
	for i := 0; i < 75; i++ {
		if i < progress {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %.0f%%", percent*100)
}

func (d *Download) GetZipPath() string {
	return d.fileZipPath
}
