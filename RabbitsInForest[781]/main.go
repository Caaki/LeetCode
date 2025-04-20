package main

func main() {}

func numRabbits(answer []int) int {

	rabbits := make(map[int]int, 0)
	for _, v := range answer {
		rabbits[v]++
	}
	result := 0
	for k, v := range rabbits {

		elements := k + 1
		if v%elements == 0 {
			result += (v / elements) * elements
		} else {
			result += (v/elements)*elements + elements
		}
	}

	return result

}
