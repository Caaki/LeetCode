package main

func maxTotalFruits(fruits [][]int, startPosInt int, k int ) int{
	result :=0

	


	if k <=2 {
		temp:=0
		for i:=0; i < k; i++{
			temp+=fruits[startPosInt-i][0]
		}
		if temp > result {
			result = temp
		}
		temp=0
		for i:=0; i < k; i++{
			temp+=fruits[startPosInt+i][0]
		}
		if temp > result {
			result = temp
		}
	}

	temp := 0
	for i:=0; i<k;i++{
		temp+=fruits[startPosInt-i][0]
	}
	if temp > result {
		result = temp
	}

	l:=k
	r := 1
	
	for l >= r {
		temp-=fruits[startPosInt-(l)][0]
		temp-=fruits[startPosInt-(l-1)][0]
		temp+=fruits[startPosInt+r][0]
		if temp > result{
			result = temp
		}
		l-=2
		r++
	} 

	for r <= k {
		temp-=fruits[l][0]
		temp+=fruits[r][0]
		temp+=fruits[r+1][0]
		if temp > result {
			result = temp
		}
		l--
		r+=2
	}

	return result

}
