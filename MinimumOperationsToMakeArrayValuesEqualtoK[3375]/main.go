package main

func main() {}

func minOperations(nums []int, k int) int {
	m := make(map[int8]bool, 0)
	for _, v := range nums {
		if v > k {
			m[int8(k)] = true
		}
		if v < k {
			return -1
		}
	}

	return len(m)
}
