package main

import (
	"fmt"
)

func main() {
	fmt.Print(smallestNumber("IDIDII"))
}

type Stack struct {
	stack []int
}

func (stack *Stack) Pop() int {
	if len(stack.stack) < 1 {
		return -1
	} else {
		temp := stack.stack[len(stack.stack)-1]
		stack.stack = stack.stack[:len(stack.stack)-1]
		return temp
	}
}

func (stack *Stack) Push(val int) {
	stack.stack = append(stack.stack, val)
}

func smallestNumber(pattern string) string {
	minVal := 1
	returnStac := make([]byte, 0, len(pattern))
	tempStack := Stack{[]int{}}
	for i := 0; i <= len(pattern); i++ {
		tempStack.Push(minVal)
		if len(pattern) == i || pattern[i] == 'I' {
			for len(tempStack.stack) > 0 {
				returnStac = append(returnStac, byte(tempStack.Pop()))
			}
		}
		minVal++
	}
	return fmt.Sprintf("%s", string(returnStac))
}

