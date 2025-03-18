package main

import(
    "fmt"
)

func main(){
    fmt.Println(longestNiceSubarray([]int{1,3,8,48,10}))
}

func longestNiceSubarray(nums []int) int {
    max:=1
    temp:=1
    win:= nums[0]
    
    for l,r:=0,1; r<len(nums); {
        if win&nums[r]==0{
            temp++
            
            win=win | nums[r]
            if max<temp{
                max = temp
            }
            r++
        }else{
            temp--
            win=win ^ nums[l]
            l++
        }
    }

    return max
}
