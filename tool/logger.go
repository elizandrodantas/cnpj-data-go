package tool

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type logger struct{}

func Logger() *logger {
	return &logger{}
}

func (l *logger) Info(at string) {
	s := color.New(color.FgYellow).Sprint("Info:")
	time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(s, time, at)
}

func (l *logger) Error(er string) {
	s := color.New(color.FgRed).Sprint("Error:")
	time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(s, time, er)
}

func (l *logger) Success(su string) {
	s := color.New(color.FgGreen).Sprint("Success:")
	time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(s, time, su)
}
