package main

func main() {
	topKFrequent([]int{-1, -1}, 1)
}

func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	valuesAndCount := make(map[int]int, 0)
	for _, n := range nums {
		valuesAndCount[n]++
	}
	count := make([][]int, len(nums)+1)

	for k, v := range valuesAndCount {
		count[v] = append(count[v], k)
	}

	var returnValues []int
	for i := len(count) - 1; i >= 0; i-- {
		if len(count[i]) > 0 {
			for _, num := range count[i] {
				returnValues = append(returnValues, num)
				k--
				if k == 0 {
					return returnValues
				}
			}
		}
	}
	return returnValues

}
