package main 


func main (){}

func maximumSum(nums []int) int{
  values := make(map[int][2]int,0)
  for i,v := range nums {
    sum :=0
    for v >0{
      sum +=v%10
      v/=10
    }
    if _,ok := values[sum]; !ok{
      values[sum] = [2]int{nums[i],-1}
      continue
    }else{
      values[sum] = getBigest2(values[sum][0],values[sum][1],nums[i])
    }
  }
  max :=-1
  for _,v:= range values{
    if v[0] ==-1 || v[1]==-1{
        continue
    }
    temp := v[0]+v[1]
    if temp > max {
      max = temp
    }
  }
  return max
}

func getBigest2(a,b,c int) [2]int{
  if a >= c && b >= c{
    return [2]int{a,b}
  }
  if b>=a && c >=a {
    return [2]int{c,b}
  }else{
    return [2]int{a,c}
  }
}
