package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func (l *Logger) Info(i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgYellow).Print(PREFIX_INFO)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Println(i...)
}

func (l *Logger) Infof(format string, i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgYellow).Print(PREFIX_INFO)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Printf(format, i...)
}
