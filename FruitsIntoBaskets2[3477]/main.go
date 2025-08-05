package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {

	count :=len(fruits)
	for _,v := range fruits {
		for j,b := range baskets {
			if v<=b{
				count--
				baskets[j]=-1
				break
			}
		}
	}

	return count
}
