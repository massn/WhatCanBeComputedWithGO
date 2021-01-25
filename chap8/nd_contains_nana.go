package chap8

import "sync"

func ndContainsNANA(input string) bool {
	targetStrings := []string{"CACA", "GAGA", "TATA", "AAAA"}
	stopCh := make(chan struct{})
	wg := &sync.WaitGroup{}
	resultCh := make(chan bool)
	defer func() { close(resultCh) }()

	for _, target := range targetStrings {
		wg.Add(1)
		go findString(target, input, resultCh, stopCh, wg)
	}
	result := false
	for i := 0; i < len(targetStrings); i++ {
		if result = <-resultCh; result {
			break
		}
	}
	close(stopCh)
	wg.Wait()
	return result
}

func findString(target, input string, resultCh chan bool, stopCh chan struct{}, wg *sync.WaitGroup) {
	defer func() { wg.Done() }()
	length := len(target)
	for i := 0; i <= len(input)-length; i++ {
		if compareSlice(target, input[i:i+length]) {
			resultCh <- true
			return
		}
		select {
		case <-stopCh:
			return
		default:
		}
	}
	resultCh <- false
}

func compareSlice(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
