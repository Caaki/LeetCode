package main

import (
	"fmt"
  "strconv"
  "math"
)


func main (){

  fmt.Println(isPalindrome2(88888))
}

func isPalindrome(x int) bool{
  strValue:= strconv.FormatInt(int64(x),10)
  j:= len(strValue)-1
  for i := 0 ; i<= len(strValue)/2;i++{
    if strValue[i] != strValue[j]{
      return false
    }
    j--;
  }
return true
}

func isPalindrome2(x int) bool{
  if x < 0 {
    return false
  }
  if x < 10 {
    return true
  }
  var s int32= int32(math.Pow(float64(10), float64(len(strconv.FormatInt(int64(x),10)) -1)))
  var m int32=0
  var l int32=0
  for x >=10{

    m = int32(x)/s
    l = int32(x)%10
    if l != m {
      fmt.Println("PreIspada",x,m,l)
      return false
    }

    x -= int(m*s)
    if int32(x)< s/10{
        if (x %100-(x%10))/10!=0{
            return false
        }

    x+=int(s/10)
    x=x/10
    x+=1
    s/=100
    fmt.Println(x)
    }else{
    x=x/10

    s /=100
    }
    
  }
  
  return true

}

