package main

import "slices"

func minimumPushes (word string) int {

	values := make(map[rune]int)

	for _,v := range word {
		values[v]+=1
	}

	counts := make([]int,0)

	for _,v := range values{
		counts = append(counts, v)
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	result :=0

	for i,v := range counts{
		if i <= 7 {
			result+=v
		}else if i <= 15{
			result += 2*v
		}else if i <= 23{
			result += 3*v
		}else {
			result += 4*v
		}
	}

	return result
}
