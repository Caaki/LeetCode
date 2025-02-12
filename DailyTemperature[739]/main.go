package main

import "fmt"

func main(){
  fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}

func dailyTemperatures(temperatures []int) []int{
  stack := Stack{[][2]int{}}
  returnVals := make([]int,len(temperatures))
  for i,v := 0,0; i < len(temperatures) ; i++{
    v = temperatures[i]
    if stack.len() < 1{
      stack.Push(v,i)
      continue
    }
    if val,index :=stack.Peek(); val < v {
      returnVals[index] = i - index
      stack.Pop()
      i--
      continue
    }
    stack.Push(v,i)
    continue
  }

  return returnVals
}

type Stack struct{
  stack [][2]int
}

func (stack *Stack) Push(v int, i int){
  if len(stack.stack) < 1{
    stack.stack = [][2]int{{v,i}}
  }else{
    stack.stack = append(stack.stack,[2]int{v,i})
  }
}

func (stack *Stack) Pop() (v int, i int){
  if len(stack.stack) < 1{
    return v,i
  }else{
    v = stack.stack[len(stack.stack)-1][0]
    i = stack.stack[len(stack.stack)-1][1]
    stack.stack = stack.stack[:len(stack.stack)-1]
    return v,i
  }
}

func (stack *Stack) Peek()(v int, i int){
  if len(stack.stack) < 1 {
    return 0,0
  }else{
    v = stack.stack[len(stack.stack)-1][0]
    i = stack.stack[len(stack.stack)-1][1]
    return v,i
  }
}

func (stack *Stack) len() int{
  return len(stack.stack)
}



