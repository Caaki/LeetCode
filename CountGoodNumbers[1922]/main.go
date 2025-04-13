package main

const mod = int(1e9 + 7)

func countGoodNumbers(n int64) int {

	even := int(n / 2)
	odd := even
	if n%2 == 0 {
		return int(getVal(5, even) * getVal(4, odd) % mod)
	} else {
		return int(getVal(5, even+1) * getVal(4, odd) % mod)

	}
}

func getVal(x, n int) int {
	if n == 0 {
		return 1
	}
	half := getVal(x, n/2)
	if n%2 != 0 {
		return (half * half * x) % mod
	}
	return (half * half) % mod
}
