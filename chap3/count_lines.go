package chap3

import (
	"strconv"
	"strings"
)

func countLines(s string) string {
	lines := strings.Split(s, "\n")
	return strconv.Itoa(len(lines))
}
