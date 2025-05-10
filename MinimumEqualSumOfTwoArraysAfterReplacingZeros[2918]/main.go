package main

func main() {}

func monSum(nums1 []int, nums2 []int) int64 {
	sum1 := 0
	sum2 := 0
	zeros1 := 0
	zeros2 := 0
	for _, v := range nums1 {
		if v == 0 {
			zeros1++
			sum1++
		}
		sum1 += v
	}

	for _, v := range nums2 {
		if v == 0 {
			zeros2++
			sum2++
		}
		sum2 += v
	}

	if sum1 > sum2 {
		sum1, sum2, zeros1, zeros2 = sum2, sum1, zeros2, zeros1
	}

	if sum1 == sum2 {
		return int64(sum1)
	}

	if zeros1 == 0 {
		return -1
	}

	return int64(sum2)

}
