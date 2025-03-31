package main

import (
	"sort"
)

func main() {}

type Val struct {
	start int
	end   int
}

type sortByStart []Val

func (s sortByStart) Len() int {
	return len(s)
}

func (s sortByStart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByStart) Less(i, j int) bool {
	return s[i].start < s[j].start
}

func partitionLabels(s string) []int {
	sVals := make(sortByStart, 26, 26)
	for i, _ := range sVals {
		sVals[i] = Val{start: -1, end: -1}
	}
	for i, v := range s {
		if sVals[v-'a'].start == -1 {
			sVals[v-'a'].start = i
			sVals[v-'a'].end = i
		} else {
			sVals[v-'a'].end = i
		}
	}
	answers := make([]int, 0)
	sort.Sort(sVals)
	l := -1
	r := -1
	for _, v := range sVals {
		if v.start == -1 {
			continue
		}
		if l == -1 {
			l = v.start
			r = v.end
			continue
		}
		if v.start > r {
			answers = append(answers, r-l+1)
			l = v.start
			r = v.end
			continue
		}
		if v.end > r {
			r = v.end
		}
	}
	if l != -1 {
		answers = append(answers, r-l+1)
	}
	return answers
}
