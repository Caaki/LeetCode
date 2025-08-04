package main


func searchInsert(nums []int, target int) int {

    l:=0
    r:=len(nums)-1

    if target < nums[0]{
        return 0
    }
    if target > nums[r]{
        return r+1
    }

    for l<=r {
        mid := l+(r-l)/2
        temp:=nums[mid]
        if temp == target{
            return mid
        }

        if temp < target{
            l=mid+1
        }else{
            r=mid-1
        }
    }
    return l
}
