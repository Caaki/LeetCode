package main

import (
  "strings"
  "fmt"
)

func main(){
  fmt.Println(numberOfSubstrings("abcabc"))
}

func rNum(r byte) int8{
  if r ==97{
    return 0
  }
  if r==98{
    return 1
  }
  return 2

}

func numberOfSubstrings(s string) int {
  abc:=make([]int16,4)
  sum:=0
  l:=-1
  r:=-1
  for i:=0; i < len(s);i++{
    if strings.ContainsRune("abc",rune(s[i])){
      if abc[3]==0{
        l=i
      }
      if abc[rNum(s[i])]==0{
        abc[3]+=1
      }
      abc[rNum(s[i])]+=1
    }
    if abc[3]==3{
      r=i
      break
    }
  }
  if r==-1{
    return sum
  }

  for  r<len(s){
    if abc[3]==3{
      sum+=len(s)-r
      if strings.ContainsRune("abc",rune(s[l])){
        abc[rNum(s[l])]-=1
        if abc[rNum(s[l])]==0{
          abc[3]-=1
        }
      }
      l++
      continue
    }
    if r<len(s)-1 && strings.ContainsRune("abc",rune(s[r+1])){
      if abc[rNum(s[r+1])]==0{
        abc[3]+=1
      }
      abc[rNum(s[r+1])]+=1
    }
    r++

  }
  return sum
}
