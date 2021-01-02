package chap4

func recognizeEvenLength(s string) string {
	i := 0
	for {
		if len(s) == i {
			return "yes"
		}
		i = i + 2
	}
}
