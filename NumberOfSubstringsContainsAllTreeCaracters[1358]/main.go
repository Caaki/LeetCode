package main

import (
  "strings"
  "fmt"
)

func main(){
  fmt.Println(numberOfSubstrings("abcabc"))
}


func numberOfSubstrings(s string) int {
  abc:=make(map[byte]int,0)
  sum:=0
  l:=-1
  r:=-1
  for i:=0; i < len(s);i++{
    if strings.ContainsRune("abc",rune(s[i])){
      if len(abc)==0{
        l=i
      }
      abc[s[i]]+=1
    }
    if len(abc)==3{
      r=i
      break
    }
  }
  if r==-1{
    return sum
  }
  started:=false
  for  r<len(s){
    if len(abc)==3{
      if !started{
        sum+=len(s)-(r-l)
        started=true
      }else{
        sum+=len(s)-r
      }
      if strings.ContainsRune("abc",rune(s[l])){
        abc[s[l]]-=1
        if abc[s[l]]==0{
          delete(abc,s[l])
        }
      }
      l++
      continue
    }
    if r<len(s)-1 && strings.ContainsRune("abc",rune(s[r+1])){
      abc[s[r+1]]+=1
    }
    r++

  }
  return sum
}
