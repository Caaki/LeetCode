package main

func main() {}

func isZeroArray(nums []int, queries [][]int) bool {

	n := len(nums)
	diffArr := make([]int, n+1)

	for _, v := range queries {
		l, r := v[0], v[1]
		diffArr[l] += 1
		diffArr[r+1] -= 1
	}

	prefixSum := make([]int, n)
	sum := 0

	for i := 0; i < len(prefixSum); i++ {
		sum += diffArr[i]
		prefixSum[i] = sum
	}

	for i, v := range nums {
		if v > prefixSum[i] {
			return false
		}
	}

	return true
}
