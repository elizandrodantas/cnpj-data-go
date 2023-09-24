package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func inputPrint(label string, sufix *string) {
	print(PREFIX_INPUT, " ", label)
	if sufix != nil {
		print(" [")
		color.New(color.FgYellow).Print(*sufix)
		print("]")
	}
	color.New(color.FgCyan).Print("?", " ")
}

func (l *Logger) InputString(label string, sufix *string) string {
	var output string

	inputPrint(label, sufix)
	fmt.Scan(&output)

	return output
}

func (l *Logger) InputInt64(label string, sufix *string) int64 {
	var output int64

	inputPrint(label, sufix)
	fmt.Scan(&output)
	print("\n")

	return output
}
