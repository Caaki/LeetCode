package main

type pair struct {
	l,r int
}

func predictTheWinner(nums []int) bool{
	pairs := make(map[pair]int)
	n:= len(nums)

	var predictRecursive func (int, int ) int
	predictRecursive = func(l,r int ) int {
		if l == r {
			return nums[l] 
		}
		k := pair{l,r}
		if v, ok := pairs[k]; ok{
			return v
		}else {

			left := nums[l] - predictRecursive(l+1, r)
			right := nums[r] - predictRecursive(l , r-1)

			if left>right{
				pairs[k]= left
				return left
			}
				pairs[k]= right
				return right
		}
	}

	return predictRecursive(0, n-1) >= 0
}
