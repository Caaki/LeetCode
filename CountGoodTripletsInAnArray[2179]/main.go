package main

func main() {}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	sortedArr := make([]int, len(nums2))

	for i, v := range nums2 {
		sortedArr[v] = i
	}

	result := int64(0)

	for i := 0; i < len(nums1)-2; i++ {
		if sortedArr[nums1[i]] > len(nums2)-3 {
			continue
		}
		pos2x := sortedArr[nums1[i]]
		for j := i + 1; j < len(nums1)-1; j++ {
			if sortedArr[nums1[j]] > len(nums2)-2 || sortedArr[nums1[j]] < pos2x {
				continue
			}
			pos2y := sortedArr[nums1[j]]
			for k := j + 1; k < len(nums1); k++ {
				if sortedArr[nums1[k]] < pos2y {
					continue
				}
				result++
			}
		}
	}

	return result

}
