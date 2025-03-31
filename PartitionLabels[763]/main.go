package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(partitionLabels("ababcbacadefegdehijhklij"))
}

type Val struct {
	start int16
	end   int16
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
			sVals[v-'a'].start = int16(i)
			sVals[v-'a'].end = int16(i)
		} else {
			sVals[v-'a'].end = int16(i)
		}
	}
	answers := make([]int, 0)
	sort.Sort(sVals)
	l := int16(-1)
	r := int16(-1)
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
			answers = append(answers, int(r-l+1))
			l = v.start
			r = v.end
			continue
		}
		if v.end > r {
			r = v.end
		}
	}
	if l != -1 {
		answers = append(answers, int(r-l+1))
	}
	return answers
}
