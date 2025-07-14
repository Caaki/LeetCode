package main

import (
	"slices"
)




func matchPlayersAndTrainers(players []int, trainers []int) int {
 	
	count := 0
		
	slices.Sort(players)
	slices.Reverse(players)
	slices.Sort(trainers)
	slices.Reverse(trainers)
	
	for i,j :=0,0 ; i<len(players) && j< len(trainers);{
		if players[i]<=trainers[j]{
			count++
			i++
			j++
		}else{
			i++
		}
	}

	return count
}
