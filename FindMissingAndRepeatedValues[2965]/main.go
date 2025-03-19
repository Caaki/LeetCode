package main

import (
  "fmt"
)

func main (){
  fmt.Println(findMissingAndRepeatedValues([][]int{{9,1,7},{8,9,2},{3,4,6}}))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
  ng:=make([]bool, len(grid)*len(grid)+1)
  returnVal:=make([]int,2,2)
  for i:=0; i<len(grid);i++{
    for j:=0;j<len(grid[i]);j++{
      if ng[grid[i][j]]{
        returnVal[0] =grid[i][j]
        continue
      }
      ng[grid[i][j]] = true
    }
  }

  for i:=1; i<len(ng);i++{
    if ng[i]==false{
      returnVal[1]= i
      break
    }
  }
  return returnVal
}
