package main

import (
	"slices"
	"sort"
)

type SaI struct {
	Val int
	I   int
}

type ValSort []SaI

func (vs ValSort) Len() int {
	return len(vs)
}

func (vs ValSort) Less(i, j int) bool {
	return vs[i].Val > vs[j].Val
}

func (vs ValSort) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

type IndexSort []*SaI

func (is IndexSort) Len() int {
	return len(is)
}

func (is IndexSort) Less(i, j int) bool {
	return is[i].I < is[j].I
}

func (is IndexSort) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func maxSubsequence(nums []int, k int) []int {
	vals := make([]SaI, 0)

	for i, v := range nums {
		vals = append(vals, SaI{Val: v, I: i})
	}

	vSort := ValSort(vals)
	sort.Sort(vSort)

	var iSort IndexSort

	for i := 0; k > 0; i++ {
		iSort = append(iSort, &vSort[i])
		k--
	}

	sort.Sort(iSort)

	toReturn := make([]int, k)
	for _, v := range iSort {
		toReturn = append(toReturn, v.Val)
	}

	return toReturn
}
