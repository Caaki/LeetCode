package main 
  
import (
  "fmt"
)

func main2 (){
  fmt.Println(minZeroArray2([]int{2,0,2}, [][]int {{0,2,1},{0,2,1},{1,1,3}}))

}


func minZeroArray2(nums []int, queries [][]int) int {
  sumOfAll:=0
  same:=true
  val:=nums[0]
  for _,v := range nums {
    if v!=val{
      same=false
    }
    sumOfAll+= v
  } 
  if sumOfAll == 0{
    return 0
  }

  it:=0
  if same{
    for i:=0; i < len(queries); i++{
      sumOfAll-=(queries[i][1]-queries[i][0]+1)*queries[i][2]
      it++
      if sumOfAll <=0{
        return it
      }
    }
  }else{
    for i:=0; i < len(queries); i++{
      decrese(nums,queries[i],&sumOfAll)
      it++
      if sumOfAll <=0{
        return it
      }
    }
  }
  return -1
}

func decrese (nums []int, q[]int, sumOfAll *int){
  for i:=q[0]; i <= q[1] ; i++{
    if nums[i]-q[2] <=0{
      *sumOfAll-= nums[i]
      nums[i]=0
    }else{
      nums[i]-=q[2]
      *sumOfAll-=q[2]
    }
  }
}
