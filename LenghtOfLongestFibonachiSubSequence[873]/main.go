package main

func main() {}

func lenLongestFibSubSeq(arr []int) int {
	longest := 0
	allValues := map[int]bool{}
	for _, v := range arr {
		allValues[v] = true
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			prev := arr[i]
			next := arr[j]
			currentCount := 2
			for {
				sum := prev + next
				_, ok := allValues[sum]
				if !ok {
					break
				}
				currentCount++
				prev = next
				next = sum
			}

			if currentCount > longest {
				longest = currentCount
			}
		}
	}
	if longest < 3 {
		return 0
	}
	return longest
}
