package main

import (
  "fmt"
)
func main(){

  fmt.Println(numTilePossibilities("AAABBC"))
}

func numTilePossibilities(tiles string) int {
  caracters:= make(map[rune]int,0)

  for _,v := range tiles{
    caracters[v]++
  }

  return getCombinationNumbers(caracters) 
}

func getCombinationNumbers(caracters map[rune]int) int{
  result:=0
  for k,v:= range caracters{
    if v > 0{
      caracters[k]--
      result++
      result+=getCombinationNumbers(caracters)
      caracters[k]++
    }
  }
  return result
}


