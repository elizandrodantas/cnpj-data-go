package logger

import (
	"fmt"
	"time"
)

func (l *Logger) Debug(t ...interface{}) {
	if !l.DebugLogger {
		return
	}

	now := time.Now().Format(TIME_FORMAT)

	print(now, " - ", PREFIX_DEBUG, " ")
	fmt.Println(t...)
}

func (l *Logger) Debugf(format string, t ...interface{}) {
	if !l.DebugLogger {
		return
	}

	now := time.Now().Format(TIME_FORMAT)

	print(now, " - ", PREFIX_DEBUG, " ")
	fmt.Printf(format, t...)
}
