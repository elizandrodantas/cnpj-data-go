package metrics

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/go-echarts/statsview"
)

const (
	PORT     = 18066
	HOST     = "localhost"
	PROTOCOL = "http"
)

func Start() {
	mgr := statsview.New()

	go mgr.Start()

	openLinkInBrowser(
		fmt.Sprintf("%s://%s:%d/debug/statsview", PROTOCOL, HOST, PORT),
	)
}

func openLinkInBrowser(link string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", link)
	case "darwin":
		cmd = exec.Command("open", link)
	default:
		cmd = exec.Command("xdg-open", link)
	}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
