package main

func kidsWithCandies(candies []int, extraCandies int) []bool{

	greatestAmount := candies[0]
	
	for _,v := range candies {
		if v > greatestAmount{
			greatestAmount = v
		}
	}
	result := make([]bool,len(candies))
	for i,v := range candies {
		if v+extraCandies >= greatestAmount {
			result[i] = true
		}
	}
	
	return result
}
