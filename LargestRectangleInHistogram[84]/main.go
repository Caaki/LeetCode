package main

import "fmt"

func main() {}

type PosVal struct {
	Pos int
	Val int
}

func largestRectangleArea(heights []int) int {
	maxVal := 0
	pStack := PosValStack{values: []PosVal{{Pos: 0, Val: heights[0]}}}
	pStack.Push(PosVal{0, heights[0]})
	i := 1
	element := pStack.Peek()
	for i < len(heights) {
		if pStack.Len() <= 0 || element.Val < heights[i] {
			pStack.Push(PosVal{Pos: i, Val: heights[i]})
			i++
			element = pStack.Peek()
			continue
		}
		if element.Val == heights[i] {
			i++
			continue
		}
		lastPosition := element.Pos
		for pStack.Len() > 0 && element.Val >= heights[i] {
			fmt.Println(pStack, element, heights[i], i)
			temp := element.Val * (i - element.Pos)
			if temp > maxVal {
				maxVal = temp
			}
			pStack.Pop()
			element = pStack.Peek()
		}
		pStack.Push(PosVal{Pos: lastPosition, Val: heights[i]})
		element = pStack.Peek()

	}

	for pStack.Len() > 0 {
		element = pStack.Peek()
		temp := element.Val * (len(heights) - element.Pos)
		if temp > maxVal {
			maxVal = temp
		}
		pStack.Pop()
	}

	return maxVal
}

type PosValStack struct {
	values []PosVal
}

func (p *PosValStack) Push(v PosVal) {
	p.values = append(p.values, v)
}

func (p *PosValStack) Len() int {
	return len(p.values)
}

func (p *PosValStack) Peek() *PosVal {
	if p.Len() <= 0 {
		return nil
	}
	return &p.values[p.Len()-1]
}

func (p *PosValStack) Pop() {
	p.values = p.values[:p.Len()-1]
}
