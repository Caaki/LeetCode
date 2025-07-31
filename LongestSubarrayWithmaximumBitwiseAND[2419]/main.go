package main

func longestSubarray(nums []int) int {
  current := nums[0]
  clen:=1
  max := nums[0]
  maxLen:=1
  for i :=1; i < len(nums);i++ {
    if nums[i] == current {
        clen++
        if max == current{
            if clen > maxLen{
                maxLen=clen
            }
        }
        if max < current {
            max = current
            maxLen=clen
        }

    }else {
        clen=1
        current = nums[i]
        if max < current {
            max = current
            maxLen=clen
        }
    }
  } 
  return maxLen
}
