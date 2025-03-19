package main

import (
  "fmt"
)

func main (){
  
  fmt.Println(coloredCells(8))
}
func coloredCells(n int) int64{
  return int64((n*4)*((n*4)-4)/4/2+1)
}



