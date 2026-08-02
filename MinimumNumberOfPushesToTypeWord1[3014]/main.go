package main

import (
	"slices"
)

func minimumPushes (word string) int {
	
	values := make(map[rune]int8)

	for _,v := range word {
		values[v]+=1
	}
	
	counts := make([]int8,0)

	for _,v := range values{
		counts = append(counts, v)
	}
	
	slices.Sort(counts)

	result :=0

	for i,v := range counts{
		if i <= 7 {
			result+=int(v)
		}else if i <= 15{
			result += 2*int(v)
		}else if i <= 23{
			result += 3*int(v)
		}else {
			result += 4*int(v)
		}
	}
	
	return result
}
