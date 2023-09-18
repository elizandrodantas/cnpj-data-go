package progress

import (
	"github.com/schollz/progressbar/v3"
)

func NewProgressDefault(text string, length int64) *progressbar.ProgressBar {
	pb := progressbar.Default(length, text)

	return pb
}

func NewProgressByte(text string, length int64) *progressbar.ProgressBar {
	pb := progressbar.DefaultBytes(length, text)

	return pb
}
