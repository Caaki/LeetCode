package main

func main() {}

type PosVal struct {
	Pos int
	Val int
}

func largestRectangleArea(heights []int) int {
	maxVal := heights[0]

	stack := Stack{s: []PosVal{{Val: maxVal, Pos: 0}}}
	current := stack.Peek()
	for i := 1; i < len(heights); {
		if current.Val == heights[i] {
			i++
			continue
		}
		if current.Val < heights[i] {
			current = PosVal{Pos: i, Val: heights[i]}

			stack.Push(current)
			i++
			continue
		}
		lastPos := current.Pos
		for stack.Len() > 0 && current.Val > heights[i] {
			lastPos = current.Pos
			temp := current.Val * (i - current.Pos)
			if maxVal < temp {
				maxVal = temp
			}
			stack.Pop()
			current = stack.Peek()
		}
		stack.Push(PosVal{Pos: lastPos, Val: heights[i]})
		current = stack.Peek()
		i++

	}

	for current.Pos != -1 {
		temp := current.Val * (len(heights) - current.Pos)
		if maxVal < temp {
			maxVal = temp
		}
		current = stack.Pop()

	}

	return maxVal

}

type Stack struct {
	s []PosVal
}

func (s *Stack) Len() int {
	return len(s.s)
}

func (s *Stack) Push(v PosVal) {
	s.s = append(s.s, v)
}
func (s *Stack) Pop() PosVal {
	if s.Len() == 0 {
		return PosVal{-1, -1}
	}
	val := s.s[s.Len()-1]
	s.s = s.s[:s.Len()-1]
	return val
}

func (s *Stack) Peek() PosVal {
	if s.Len() == 0 {
		return PosVal{-1, -1}
	}
	return s.s[s.Len()-1]
}
