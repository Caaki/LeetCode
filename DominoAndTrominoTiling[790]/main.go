package main

func main() {}

func numTilings(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if n == 3 {
		return 5
	}
	mod := 1000000007

	answers := make([]int, n+1)

	answers[1] = 1
	answers[2] = 2
	answers[3] = 5
	var myFunc func(int, int) int
	myFunc = func(m int, prev int) int {
		if m == n {
			return (prev*2 + answers[m-3]) % mod
		}
		nVal := (prev*2 + answers[m-3]) % mod
		answers[m] = nVal
		return myFunc(m+1, nVal)
	}

	return myFunc(4, 5)
}
