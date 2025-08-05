package main

func totalFruit(fruits []int) int{

	l:=0
	r:=0

	values := make(map[int]int,0)
	maxLen:=0
	
	for r<len(fruits){
		values[fruits[r]]++
	
		for len(values) > 2{
			values[fruits[l]]--
			if values[fruits[l]]==0{
				delete(values, fruits[l])
			}
			l++
		}

		if maxLen < r+1-l  {
			maxLen=r+1-l
		}

		r++
	}


	return maxLen
}
