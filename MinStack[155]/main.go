package main

func main() {}

type MinStack struct {
	stack []int
	Min   []int16
}

func Constructor() MinStack {
	minStack := MinStack{[]int{}, []int16{}}
	return minStack
}

func (this *MinStack) Push(val int) {
	if len(this.stack) <= 0 {
		this.stack = append(this.stack, val)
		this.Min = append(this.Min, 0)
	} else {
		this.stack = append(this.stack, val)
		if this.stack[this.Min[len(this.Min)-1]] >= val {
			this.Min = append(this.Min, int16(len(this.stack)-1))
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
	} else {
		if len(this.stack)-1 == int(this.Min[len(this.Min)-1]) {
			this.Min = this.Min[:len(this.Min)-1]
			this.stack = this.stack[:len(this.stack)-1]
		} else {
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) <= 0 {
		return 0
	} else {
		return this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) GetMin() int {
	if len(this.stack) <= 0 {
		return 0
	} else {
		return this.stack[this.Min[len(this.Min)-1]]
	}
}
