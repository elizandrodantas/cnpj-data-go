package migration

import (
	"fmt"
	"regexp"

	"golang.org/x/text/encoding/charmap"
)

var replaceStringFunc = func(s string) string {
	switch s {
	case "'":
		return "`"
	case ";":
		return ","
	case "\"":
		return "'"
	case "\\":
		return "/"
	default:
		return s
	}
}

func parseInsertValues(d []string) []string {
	var output []string

	regex := regexp.MustCompile(`[';"\\]`)

	for _, k := range d {
		parse := fmt.Sprintf("(%s)", regex.ReplaceAllStringFunc(k, replaceStringFunc))
		utf8, _ := charmap.ISO8859_1.NewDecoder().String(parse)

		output = append(output, utf8)
	}

	return output
}

func parseNil(d []string) []string {
	var output []string

	for _, v := range d {
		if len(v) > 1 {
			byt := []byte{}

			for _, b := range []byte(v) {
				if b != 0 {
					byt = append(byt, b)
				}
			}

			output = append(output, string(byt))
		}
	}

	return output
}
