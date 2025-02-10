package main

import (
	"strconv"
)

func main(){}


func evalRPN(tokens []string) int{
  if len(tokens) == 1{
    val,_ := strconv.Atoi(tokens[0])
    return val
  }
  n:= Stack{}
  for _,val:= range tokens{
    if val != "+" && val!="-" && val!= "*" && val!= "/" {
      num ,_ := strconv.Atoi(val)
      n.Push(num)
    }else{
        n.Push(calc(n.Pop(),n.Pop(),val))
    }
  }
  return n.Pop()
}

func calc(a int, b int, c string) int{
  switch c {
  case ("+"):
    return b+a
  case ("-"):
    return b-a
  case("*"):
    return b*a
  default:
    return b/a
  }
}

type Stack struct{
  stack []int
}

func (stack *Stack) Push (val int) {
  if len(stack.stack) <=0{
    stack.stack = []int{val}
  }else{
    stack.stack = append(stack.stack, val)
  }
}

func (stack *Stack) Pop () int{
  if len(stack.stack)<= 0{
    return 0
  }else{
    element := stack.stack[len(stack.stack)-1]
    stack.stack = stack.stack[:len(stack.stack)-1]
    return element
  }
}
