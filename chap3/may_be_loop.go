package chap3

import (
	"strings"
)

const keyword = "secret sauce"

func mayBeLoop(s string) string {
	if strings.Contains(s, keyword) {
		if len(s)%2 == 0 {
			return "yes"
		} else {
			return "no"
		}
	} else {
		i := 0
		for {
			i += 1
		}
	}
}
