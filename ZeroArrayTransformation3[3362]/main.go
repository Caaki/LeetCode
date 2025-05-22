package main

import "sort"

func main() {}

func maxRemoval(nums []int, queries [][]int) int {

	sort.Sort(Queries(queries))

	n := len(nums)

	diffArr := make([]int, n+1)

	for _, v := range queries {
		l, r := v[0], v[1]
		diffArr[l] += 1
		diffArr[r+1] -= 1
	}

	prefixSum := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		sum += diffArr[i]
		prefixSum[i] = sum
	}

	for i, v := range nums {
		if v > prefixSum[i] {
			return -1
		}
	}

	result := 0
	for _, v := range queries {
		l, r := v[0], v[1]
		copyDiff := make([]int, len(diffArr))
		copy(copyDiff, diffArr)
		copyDiff[l] -= 1
		copyDiff[r+1] += 1
		pSum := make([]int, n)
		sum = 0

		for i := 0; i < n; i++ {
			sum += copyDiff[i]
			pSum[i] = sum
		}

		test := true
		for i, v := range nums {
			if v > pSum[i] {
				test = false
				break
			}
		}

		if test {
			result++
			copy(diffArr, copyDiff)
		}
	}
	return result
}

type Queries [][]int

func (q Queries) Len() int {
	return len(q)
}

func (q Queries) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q Queries) Less(i, j int) bool {
	return (q[i][1] - q[i][0]) < (q[j][1] - q[j][0])
}
