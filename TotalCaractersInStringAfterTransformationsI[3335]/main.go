package main

import "fmt"

const mod int = 1000000007

func main() {
	fmt.Println(lengthAfterTransformations("abcyy", 2))
}

func lengthAfterTransformations(s string, t int) int {
	caracters := make([]int, 26)

	for _, v := range s {
		caracters[v-'a']++
	}

	for i := 0; i < t; i++ {
		temp := make([]int, 26)
		z := caracters[25]
		for j := 0; j <= 24; j++ {
			temp[j+1] = (temp[j+1] + caracters[j]) % mod
		}
		if z != 0 {
			temp[0] = (temp[0] + z) % mod
			temp[1] = (temp[1] + z) % mod
		}
		caracters = temp
	}

	count := 0
	for _, v := range caracters {
		count = (count + v) % mod
	}

	return count

	return 0
}
