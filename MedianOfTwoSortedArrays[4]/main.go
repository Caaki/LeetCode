package main

import(
  "fmt"
)

func main (){

  fmt.Println(findMedianSortedArrays([]int{1,3},[]int{2}))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  n:=0
  m:=0
  combinedLen:= len(nums1) + len(nums2)
  combinedArray := make([]int, combinedLen/2)
  for i := 0 ; i <= combinedLen ; i ++{
    if n >= len(nums1){
        combinedArray = append(combinedArray, nums2[m])
        m++
        continue
    }
        if m >= len(nums2){
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
