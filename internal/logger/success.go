package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func (l *Logger) Success(i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgGreen).Print(PREFIX_SUCCESS)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Println(i...)
}

func (l *Logger) Successf(format string, i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgGreen).Print(PREFIX_SUCCESS)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Printf(format, i...)
}
