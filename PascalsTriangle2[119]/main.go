package main

func getRow(rowIndex int) []int {
	result := make([][]int,rowIndex+1)

	for i,size:=0,1 ; i <= rowIndex; i++{
		result[i] = make([]int, size)
		result[i][0]=1
		result[i][len(result[i])-1]=1
        size++
	}
	
	if rowIndex <2 {
		return result[rowIndex]
	}
	
	for i := 2 ; i <= rowIndex ; i++{
		for j:= 1; j+1 < len(result[i]) ; j++{
			result[i][j] = result[i-1][j-1] + result[i-1][j] 
		}
	}
	return result[rowIndex]
}
