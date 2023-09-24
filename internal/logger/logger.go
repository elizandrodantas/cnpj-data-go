package logger

import (
	"time"
)

const (
	// TIME
	TIME_FORMAT = time.DateTime

	// INFO OPTIONS
	PREFIX_INFO = "[i]"

	// SUCCESS OPTIONS
	PREFIX_SUCCESS = "[+]"

	// ERROR OPTIONS
	PREFIX_ERROR = "[X]"

	// DEBUG OPTIONS
	PREFIX_DEBUG = "DEBUG:"

	// INPUT OPTIONS
	PREFIX_INPUT = ">"
	SUFIX_INPUT  = "?"
)

type Logger struct {
	DebugLogger bool
}
