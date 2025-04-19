package main

import (
  "sort"
)

func main() {
  countFairPairs([]int{5, 2, 3, 1, 67, 1, 3}, 1, 2)
}

func countFairPairs(nums []int, lower int, upper int) int64 {

  sort.Ints(nums)
  fmt.Println(nums)
  result:=int64(0)

  for i:=0;i < len(nums); i++{
    l,r := bSearch(nums[i:], lower-nums[i], upper-nums[i] )
    if r==-1{
      continue
    }
    result+=int64(r-l+1)
  }
  return result
}


func bSearch (nums []int, lower,upper int) (int,int){

  l:=0
  r:= len(nums) -1
  lPos:=-1
  rPos:=-1
  closest:=-1
  for l<=r {
    mid:= l + (r-l)/2
    if nums[mid] == lower{
      lPos=mid
      break
    }
    if nums[mid]<lower{
      l=mid+1
      closest=mid
    }else{
      r = mid-1
      closest=mid
    }
  } 

    fmt.Println(nums[closest])
  if lPos!=-1{
    for i:=lPos-1; i>=0 ;i--{
      if nums[i] == lower{
        lPos=i
      }
    }
  }else{
    if nums[closest] <lower{
      for i:=closest+1;i< len(nums);i++{
        if nums[closest]>=lower && nums[closest]<=upper{
          lPos=i
        }
        if nums[closest]>upper{
          break
        }
      }
    }else{
      for i:=closest;i>=0;i--{
        if nums[closest]>=lower && nums[closest]<=upper{
          lPos=i
        }
        if nums[closest]>upper{
          break
        }
      }
    }
  }
    if lPos==-1{
      return -1,-1
    }

  closest=-1
  for l<=r {
    mid:= l + (r-l)/2
    if nums[mid] == upper{
      rPos=mid
      break
    }
    if nums[mid]>upper{
      l=mid+1
      closest=mid
    }else{
      r = mid-1
      closest=mid
    }
  } 

  if rPos!=-1 {
    for i:=rPos+1; i<len(nums) ;i++{
      if nums[i] == upper{
        rPos=i
      }
    }
  }else{
    if nums[closest] >upper{
      for i:=closest-1;i>=0;i--{
        if nums[closest]>=lower && nums[closest]<=upper{
          :         rPos=i
        }
        if nums[closest] < lower{
          break
        }
      }
    }else{
        for i:=closest;i<len(nums);i++{
          if nums[closest]>=lower && nums[closest]<=upper{
            rPos=i
          }
          if nums[closest] < lower {
            break
          }
        }
      }
    }
      if rPos==-1{
        return -1,-1
      }
    return lPos,rPos


}
