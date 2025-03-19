package main

func main() {
	topKFrequent([]int{-1, -1}, 1)
}

func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	valuesAndCount := make(map[int16]int8, 0)
	for _, n := range nums {
		valuesAndCount[int16(n)]++
	}
	count := make([][]int16, len(nums)+1)

	for k, v := range valuesAndCount {
		count[v] = append(count[v], k)
	}

	var returnValues []int
	for i := len(count) - 1; i >= 0; i-- {
		if len(count[i]) > 0 {
			for _, num := range count[i] {
				returnValues = append(returnValues, int(num))
				k--
				if k == 0 {
					return returnValues
				}
			}
		}
	}
	return returnValues

}
