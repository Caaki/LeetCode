package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println(isValid("[]{}(())"))

}

type Stack[T any] struct {
	stack []T
}

func MakeStack[T any](size ...int) Stack[T] {

	if len(size) > 0 {
		return Stack[T]{
			stack: make([]T, 0, size[0]),
		}
	}

	return Stack[T]{
		stack: make([]T, 0),
	}
}

func (stack *Stack[T]) Push(element T) {
	stack.stack = append(stack.stack, element)
}

func (stack *Stack[T]) Pop() (element T) {
	stackLen := len(stack.stack) - 1
	if stackLen >= 0 {
		element = stack.stack[stackLen]
		stack.stack = slices.Delete(stack.stack, stackLen, stackLen+1)
		return element
	}
	var zeroValue T
	return zeroValue
}

func (stack *Stack[T]) Size() int {

	return len(stack.stack)
}

func isValid(s string) bool {
	stack := MakeStack[rune]()

	for _, v := range s {
		if isOpened(v) {
			stack.Push(v)
		} else {
			if stack.Size() < 0 {
				return false
			}
			if stack.Pop() != getNeededClosed(v) {
				return false
			}
		}
	}
	if stack.Size() > 0 {
		return false
	}
	return true
}
func getNeededClosed(s rune) rune {
	if s == '}' {
		return '{'
	}
	if s == ')' {
		return '('
	}
	return '['

}

func isOpened(s rune) bool {
	if s == '(' || s == '{' || s == '[' {
		return true
	}
	return false
}
