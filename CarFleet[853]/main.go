package main

import (
  "fmt"
)
func main(){
  carFleet(12,[]int{4,1,0,7},[]int{2,2,1,1})
}


func carFleet(target int, position []int, speed []int) int{
  fmt.Println(position,speed)
  quickSort(position,speed,0,len(position)-1)
  fmt.Println(position,speed)
  return 0
}

func quickSort(position []int, speed []int, start int, end int) {

  if end <= start{
    return
  }

  pivot := partition(position, speed,start,end)
  quickSort(position, speed, start, pivot-1)
  quickSort(position, speed, pivot+1, end)
}

func partition (position []int, speed []int, start int ,end int) int{
  pivot:= position[end]
  i := start -1
  temp :=0
  for j := start ; j <= end-1 ; j++{
    if position[j] < pivot{
      i++
    temp = position[i]
    position[i] = position[j]
    position[j] = temp

    temp = speed[i]
    speed[i] = speed[j]
    speed[j] = temp
    }
  }
  i++
  temp = position[i]
  position[i] = position[end]
  position[end] = temp

  temp = speed[i]
  speed[i] = speed[end]
  speed[end] = temp



  return i
}
