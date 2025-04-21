package main

func main() {}

func numberOfArrays(differences []int, lower int, upper int) int {
	current := 0
	maxinumum := 0
	minimum := 0
	for _, v := range differences {
		current += v
		if minimum > current {
			minimum = current
		}
		if maxinumum < current {
			maxinumum = current
		}
	}

	dif := upper - lower
	dif2 := maxinumum - minimum
	if dif2 > dif {
		return 0
	}

	return dif - dif2 + 1
}
