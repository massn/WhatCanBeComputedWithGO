package chap2

import (
	"strings"
)

func ContainsGAGA(s string) string {
	if strings.Contains(s, "GAGA") {
		return "yes"
	}
	return "no"
}
