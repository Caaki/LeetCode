package main

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {

	maxVal := getDiagonal(dimensions[0][0], dimensions[0][1])
	maxR:=dimensions[0][0] * dimensions[0][1]
	for _,v := range dimensions{
		temp := getDiagonal(v[0], v[1])
		temp2 := v[0]* v[1]
		if maxVal <= temp {
			if maxVal < temp{
				maxR = v[0] * v[1]
				maxVal = temp
				continue
			}

			if maxR < temp2{
				maxR = v[0] * v[1]
			}
		}
	}
	return maxR
}

func getDiagonal (i,j int ) float64 {
	return math.Sqrt(float64(i*i) +float64(j*j))
}
