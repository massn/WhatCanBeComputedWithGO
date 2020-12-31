package chap3

func longerThan1k(s string) string {
	if len(s) > 1000 {
		return "yes"
	}
	return "no"
}
