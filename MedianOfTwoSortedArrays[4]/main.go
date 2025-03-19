package main

import (
	"fmt"
	"math"
)

func main (){

  fmt.Println(findMedianSortedArrayBinarySearch([]int{1,3},[]int{2}))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  n:=0
  m:=0
  combinedLen:= len(nums1) + len(nums2)
  combinedArray := make([]int, combinedLen/2)
  for i := 0 ; i <= combinedLen ; i ++{
    if n >= len(nums1)-1{
        combinedArray = append(combinedArray, nums2[m])
        m++
        continue
    }
        if m >= len(nums2)-1{
        combinedArray = append(combinedArray, nums1[n])
        n++
        continue
    }
    if nums1[n] < nums2[m]{
      combinedArray = append(combinedArray, nums1[n])
      n++
    }else{
      combinedArray = append(combinedArray, nums2[m])
      m++
    }
  }
  if combinedLen%2 ==0 {
    return float64(combinedArray[combinedLen-1] / combinedArray[combinedLen-2] )
  }
  return float64(combinedArray[combinedLen-1])
}


func findMedianSortedArrayBinarySearch(nums1 []int, nums2[]int) float64{
  if len(nums1) > len(nums2){
    nums1, nums2 = nums2,nums1
  }
  
  x:= len(nums1)
  y := len(nums2)
  
  if x==0{
    if y%2==0{
      return (float64(nums2[y/2-1])+float64(nums2[y/2]))/2
    }
    return float64(nums2[y/2])
  }

  low := 0
  high:= x
  var maxLeftX int
  var minRightX int

  var maxLeftY int
  var minRightY int
  for low <= high{
  
    positionX:= (low+high) /2
    positionY:= (x+y+1)/2 - positionX
    
    if positionX <=0 {
      maxLeftX = math.MinInt32
    }else{
      maxLeftX = nums1[positionX-1]
    }
    
    if positionX >= x {
      minRightX = math.MaxInt
    }else{
      minRightX = nums1[positionX]
    }

    if positionY <=0{
      maxLeftY = math.MinInt32
    }else{
      maxLeftY = nums2[positionY-1]
    }
 
    if positionY >= y {
      minRightY = math.MaxInt
    }else{
      minRightY = nums2[positionY]
    }

    if maxLeftX <= minRightY && maxLeftY <= minRightX {
      if (x+y)%2==0{
        return (math.Max(float64(maxLeftX),float64(maxLeftY))+math.Min(float64(minRightX),float64(minRightY)))/2
      }
      return math.Max(float64(maxLeftX),float64(maxLeftY))
    }
    
    if maxLeftX > minRightY{
      high = positionX-1
    }else{
      low = positionX+1
    }
  }
  return 0
}








