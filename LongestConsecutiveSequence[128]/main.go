package main

import (

  "fmt"
  "time"
)

func main(){

  fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}

func longestConsecutive(nums []int) int{
  values:= make(map[int][]int)
  longest:=0
  for _,num :=range nums{
    if _,ok:= values[num];ok{
      continue
    }
    values[num] = []int{num}
    time.Sleep(1*time.Second)
    fmt.Println("num is",num,"and it has",values[num])
    getFirst(values, num,num)
    getLast(values,num)
    if longest < len(values[num]){
      longest = len(values[num])
    }
  }
  return longest
}

func getFirst(values map[int][]int, num int, appendTo int){
  if _,ok:= values[num-1];ok{
    time.Sleep(1*time.Second)

    if values[num-1][len(values[num-1])-1] == num{
  values[appendTo] = append(values[num-1][:len(values[num-1])],values[appendTo]...)
    }else{
    values[appendTo]=append(values[num-1], values[appendTo]...)
    }
    fmt.Println(values[appendTo])
    getFirst(values, values[appendTo][0],appendTo)
  }
}

func getLast(values map[int][]int, num int){
  if _,ok:= values[num+1];ok{
    time.Sleep(1*time.Second)
    values[num] = append(values[num],values[num+1]...)
    fmt.Println(values[num])
    getLast(values, values[num+1][0])
  }
}

