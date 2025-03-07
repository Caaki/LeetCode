package main

import(
  "math"
  "fmt"
)

func main (){
  fmt.Println(closestPrimes(10,19))
}

func closestPrimes(left int, right int) []int{
  if right ==3 {
    return []int{2,3}
  }
  primes:=make([]bool,right+1)
  primes[0]=true
  primes[1]=true
  to:=int(math.Sqrt(float64(right)))
  toCheck:=[]int32{}
  for i:=2; i <= to;i++{
    if !primes[i] {
      for j:= i*2; j<=right; j+=i{
        primes[j]=true
      }
    }
  }

  for i:=left;i<len(primes);i++{
    if !primes[i]{
      toCheck=append(toCheck,int32(i))
    }
  }
  ln:=len(toCheck)

  if ln < 2{
    return []int{-1,-1}
  }
  pos1:=int8(ln-1)
  pos2:=int8(ln-2)

  distance:=int(toCheck[ln-1]-toCheck[ln-2])

  for i:=ln-3; i>0;i--{
    if toCheck[i]-toCheck[i-1]<= int32(distance){
      pos1=int8(i)
      pos2=int8(i-1)
      distance=int(toCheck[i]-toCheck[i-1])
    }
  }
  return []int{int(toCheck[pos2]),int(toCheck[pos1])}
}
