package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/elizandrodantas/cnpj-data-go/internal/progress"
)

type download struct {
	name string
}

func Start(link string) (*download, error) {
	if !strings.HasPrefix(link, "http") {
		return nil, fmt.Errorf("this link is invalid")
	}

	if _, err := url.Parse(link); err != nil {
		return nil, fmt.Errorf("link parse error")
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	response, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	tempFile, err := os.CreateTemp("", "gov-br-*.zip")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	name := tempFile.Name()
	contentLength := response.ContentLength

	bp := progress.NewProgressByte("downloading", contentLength)

	_, err = io.Copy(io.MultiWriter(tempFile, bp), response.Body)
	if err != nil {
		return nil, err
	}

	return &download{name}, nil
}

func (d *download) GetName() string {
	return d.name
}
