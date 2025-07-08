package main

func main() {}

type FindSumPairs struct {
	Arr1 []int
	Arr2 []int
	MV   map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m := make(map[int]int, 0)

	for _, v := range nums2 {
		m[v]++
	}

	f := FindSumPairs{Arr1: nums1, Arr2: nums2, MV: m}
	return f
}

func (this *FindSumPairs) Add(index int, val int) {
	temp := this.Arr2[index]
	this.Arr2[index] += val
	this.MV[temp]--
	this.MV[temp+val]++
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0
	for _, v := range this.Arr1 {
		need := tot - v
		if this.MV[need] > 0 {
			count += this.MV[need]
		}
	}
	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
