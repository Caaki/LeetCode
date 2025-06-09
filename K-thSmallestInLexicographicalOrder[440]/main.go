package main

func main() {}

func findKthNumber(n int, k int) int {
	count := 0
	finished := false
	val := 0
	for j := 1; j <= 9 && j <= n; j++ {
		addValues(j, n, k, &count, &finished, &val)
		if finished {
			break
		}
	}

	return val

}

func addValues(current, n, k int, count *int, finished *bool, val *int) {
	if *finished || current > n {
		return
	}
	*count = *count + 1
	if *count == k {
		*finished = true
		*val = current
		return
	}
	for i := 0; i < 10; i++ {
		next := current*10 + i
		if next > n {
			return
		}
		addValues(next, n, k, count, finished, val)
	}
}
