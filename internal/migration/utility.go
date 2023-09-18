package migration

import (
	"fmt"
	"strings"
)

func parseInsertValues(d []string) []string {
	var output []string

	for _, k := range d {
		parse := fmt.Sprintf("(%s)", strings.ReplaceAll(k, ";", ","))

		output = append(output, parse)
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
