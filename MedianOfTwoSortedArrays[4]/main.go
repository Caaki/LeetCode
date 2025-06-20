package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(findMedianSortedArrayBinarySearch([]int{1, 3}, []int{2}))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := 0
	m := 0
	combinedLen := len(nums1) + len(nums2)
	combinedArray := make([]int, combinedLen/2)
	for i := 0; i <= combinedLen; i++ {
		if n >= len(nums1)-1 {
			combinedArray = append(combinedArray, nums2[m])
			m++
			continue
		}
		if m >= len(nums2)-1 {
			combinedArray = append(combinedArray, nums1[n])
			n++
			continue
		}
		if nums1[n] < nums2[m] {
			combinedArray = append(combinedArray, nums1[n])
			n++
		} else {
			combinedArray = append(combinedArray, nums2[m])
			m++
		}
	}
	if combinedLen%2 == 0 {
		return float64(combinedArray[combinedLen-1] / combinedArray[combinedLen-2])
	}
	return float64(combinedArray[combinedLen-1])
}

func findMedianSortedArrayBinarySearch(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	mid := 0
	l := 0
	r := len(nums1)
	h := (len(nums1) + len(nums2)) / 2
	for l <= r {
		mid = l + (r-l)/2
		x1 := math.MinInt
		y1 := math.MaxInt

		if mid > 0 {
			x1 = nums1[mid-1]
		}
		if mid < len(nums1) {
			y1 = nums1[mid]
		}

		mid2 := h - mid

		x2 := math.MinInt
		y2 := math.MaxInt

		if mid2 > 0 {
			x2 = nums2[mid2-1]
		}

		if mid2 < len(nums2) {
			y2 = nums2[mid2]
		}

		if x1 <= y2 && x2 <= y1 {
			if (len(nums1)+len(nums2))%2 == 0 {
				return (findMax(x1, x2) + findMin(y1, y2)) / 2
			} else {
				return findMin(y1, y2)
			}
		}

		if x1 > y2 {
			r = mid - 1
			continue
		}
		l = mid + 1

	}

	return 1
}

func findMax(i, j int) float64 {
	if i > j {
		return float64(i)
	} else {
		return float64(j)
	}
}
func findMin(i, j int) float64 {
	if i > j {
		return float64(j)
	} else {
		return float64(i)
	}
}
