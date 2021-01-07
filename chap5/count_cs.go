package chap5

import "fmt"

type move int
type state int
type tape []rune

const (
	right move = iota
	left
	stay
)

type Written struct {
	write     bool
	character rune
}

type TuringMachine interface {
	Reset()
	Transit(rune) (Written, move)
	IsHaltState() bool
	IsRejectState() bool
}

func Run(t TuringMachine, strTape string) (bool, string) {
	tape := stringToTape(strTape)
	for i := 0; ; {
		written, move := t.Transit(tape[i])
		if written.write {
			tape[i] = written.character
		}
		switch move {
		case right:
			i += 1
		case left:
			i -= 1
		case stay:
		default:
			panic(fmt.Sprintf("undefined move %v", move))
		}
		if t.IsHaltState() {
			return true, tapeToString(tape)
		}
		if t.IsRejectState() {
			return false, tapeToString(tape)
		}
	}
}

func reverse(str string) string {
	n := len(str)
	reversed := ""
	for i := n - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func stringToTape(str string) tape {
	tape := []rune{}
	for _, c := range str {
		tape = append(tape, rune(c))
	}
	return tape
}

func tapeToString(tape tape) string {
	str := ""
	for _, r := range tape {
		str += string(r)
	}
	return str
}

const (
	binaryIncrementerQ0 state = iota
	binaryIncrementerQhalt
	binaryIncrementerQreject
	binaryIncrementerQ1
	binaryIncrementerQ2
	binaryIncrementerQ3
)

type binaryIncrementer struct {
	State state
}

func NewBinaryIncrementer() *binaryIncrementer {
	return &binaryIncrementer{State: binaryIncrementerQ0}
}

func (t *binaryIncrementer) Reset() {
	t.State = binaryIncrementerQ0
}

func (t *binaryIncrementer) IsHaltState() bool {
	return t.State == binaryIncrementerQhalt
}

func (t *binaryIncrementer) IsRejectState() bool {
	return t.State == binaryIncrementerQreject
}

func (t *binaryIncrementer) Transit(read rune) (Written, move) {
	switch t.State {
	case binaryIncrementerQ0:
		if read == 'x' {
			t.State = binaryIncrementerQ1
			return writeNothing(), right
		}
		return t.reject()
	case binaryIncrementerQ1:
		switch read {
		case '0', '1':
			t.State = binaryIncrementerQ1
			return writeNothing(), right
		case 'x':
			t.State = binaryIncrementerQ2
			return writeNothing(), left
		default:
			return t.reject()
		}
	case binaryIncrementerQ2:
		switch read {
		case '0':
			t.State = binaryIncrementerQ3
			return writeChar('1'), left
		case '1':
			t.State = binaryIncrementerQ2
			return writeChar('0'), left
		case 'x':
			t.State = binaryIncrementerQhalt
			return writeChar('1'), stay
		default:
			return t.reject()
		}
	case binaryIncrementerQ3:
		switch read {
		case '0', '1':
			t.State = binaryIncrementerQ3
			return writeNothing(), left
		case 'x':
			t.State = binaryIncrementerQhalt
			return writeNothing(), stay
		default:
			return t.reject()
		}
	default:
		return t.reject()
	}
}

func (t *binaryIncrementer) reject() (Written, move) {
	t.State = binaryIncrementerQreject
	return Written{write: false}, stay
}
func writeNothing() Written {
	return Written{write: false}
}

func writeChar(c rune) Written {
	return Written{write: true, character: c}
}
