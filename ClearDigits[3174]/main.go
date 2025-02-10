package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Print(clearDigits("qm93xjkpaaovhqckjhg5j1o4rndtg3bobgeke"))
}

func clearDigits(s string) string {
	toPrint := Stack{}
	for _, val := range s {
		if unicode.IsDigit(val) {
			toPrint.Pop()
		} else {
			toPrint.Push(byte(val))
		}
	}
	return string(toPrint.stack)
}

type Stack struct {
	stack []byte
}

func (stack *Stack) Push(val byte) {
	if len(stack.stack) <= 0 {
		stack.stack = []byte{val}
	} else {
		stack.stack = append(stack.stack, val)
	}
}

func (stack *Stack) Pop() {
	if len(stack.stack) <= 0 {
	} else {
		stack.stack = stack.stack[:len(stack.stack)-1]
	}
}
