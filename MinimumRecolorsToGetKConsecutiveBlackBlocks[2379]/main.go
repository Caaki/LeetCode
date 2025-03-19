package main

import(
  "fmt"
)
func main(){
  fmt.Println(minimumRecolors("WWBBBWBBBBBWWBWWWB",16))
}


func minimumRecolors(blocks string, k int) int{
  mc:=0
  temp:=0
  for i:=0; i < k; i++{
    if blocks[i] =='W'{
      temp++ 
    }
  }
  mc = temp
  if mc ==0 {
    return 0
  }
  
  for i:=k; i<len(blocks); i++{
    if blocks[i-k] == blocks[i]{
      continue
    }
    
    if blocks[i] =='W'{
      temp++
    }else{
      temp--
    }
    
    if mc > temp{
      mc=temp
    }
    if mc <=0{
        return 0
    }
  }
  
  return mc
}
