package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func (l *Logger) Error(i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgRed).Print(PREFIX_ERROR)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Println(i...)
}

func (l *Logger) Errorf(format string, i ...interface{}) {
	now := time.Now().Format(TIME_FORMAT)

	color.New(color.FgRed).Print(PREFIX_ERROR)
	color.New(color.FgMagenta).Print(" [", now, "] ")
	fmt.Printf(format, i...)
}
