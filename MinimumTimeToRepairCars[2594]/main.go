package main
import (
  "math"
  "fmt"
)

func main(){
  fmt.Println(repairCars([]int{4,2,3,1},10))
}


func repairCars(ranks []int, cars int) int64 {
  l:=1

  h:=ranks[0]*(cars*cars)

  returnVal:=-1
  for l<=h{
    mid := (l+h)/2
    sum:=0
    for _,v:=range ranks{
      sum+=int(math.Sqrt(float64(mid)/float64(v)))
    }
    if sum < cars{
      l=mid+1
    }else{
      returnVal=mid
      h=mid-1
    }
  }

  return int64(returnVal)
}
