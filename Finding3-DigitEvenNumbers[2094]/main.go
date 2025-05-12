package main

func main() {}

func findEvenNUmbers(digits []int) []int {

	answer := make([]int, 0)
	values := make([]int, 10)

	for _, v := range digits {
		values[v]++
	}

	for i := 1; i <= 9; i++ {
		if values[i] == 0 {
			continue
		}
		for j := 0; j <= 9; j++ {
			if values[j] == 0 || (i == j && values[i] < 2) {
				continue
			}
			for k := 0; k <= 8; k += 2 {
				if values[k] == 0 {
					continue
				}
				count := 1
				if k == i {
					count++
				}
				if k == j {
					count++
				}
				if count > values[k] {
					continue
				}
				answer = append(answer, i*100+j*10+k)
			}
		}
	}

	return answer
}
