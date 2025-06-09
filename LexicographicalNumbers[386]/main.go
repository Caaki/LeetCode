package main

func main() {}

func lexicalOrder(n int) []int {

	values := make([]int, 0)

	for j := 1; j <= 9 && j <= n; j++ {
		values = append(values, j)
		addValues(j, n, &values)
	}
	return values
}

func addValues(current, to int, arr *[]int) {
	for i := 0; i < 10; i++ {
		next := current*10 + i
		if next > to {
			return
		}
		*arr = append(*arr, next)
		addValues(next, to, arr)
	}
}
