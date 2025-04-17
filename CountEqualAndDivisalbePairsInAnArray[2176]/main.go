package main

func main() {}

func countPairs(nums []int, k int) int {

	values := make([][]int16, 101)

	for i, v := range nums {
		values[v] = append(values[v], int16(i))
	}

	result := 0
	for _, arr := range values {
		if len(arr) > 1 {
			for i := 0; i < len(arr); i++ {
				for j := i + 1; j < len(arr); j++ {
					if int(arr[i]*arr[j])%k == 0 {
						result++
					}
				}
			}
		}
	}

	return result
}
