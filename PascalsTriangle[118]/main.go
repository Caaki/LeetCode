package main

func generate (numRows int) [][]int {
	result := make([][]int,numRows)

	for i,size:=0,1 ; i < numRows; i++{
		result[i] = make([]int, size)
		result[i][0]=1
		result[i][len(result[i])-1]=1
		size++
	}
	
	if numRows <=2 {
		return result
	}
	
	for i := 2 ; i < numRows ; i++{
		for j:= 1; j+1 < len(result[i]) ; j++{
			result[i][j] = result[i-1][j-1] + result[i-1][j] 
		}
	}
	return result
}
