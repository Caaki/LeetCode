package main

import "fmt"

func main (){
  fmt.Println(lenghtOfLongestSubstring("dvdf"))
}


func lenghtOfLongestSubstring(s string) int {

  current := make(map[rune]int16,len(s)/50)
  maxCount := 0
  currentCount :=0

  for i := int16(0) ; i < int16(len(s)) ; i++{
    if _,ok := current[rune(s[i])] ; ok{  
      if maxCount < currentCount{
        maxCount = currentCount
      } 
      currentCount = 1
      i = current[rune(s[i])]+1
      clear(current)
      current[rune(s[i])] = i
    }else{
        current[rune(s[i])]=i
        currentCount ++
    }
  }
  if maxCount<currentCount{
    return currentCount
  }
  return maxCount

}
