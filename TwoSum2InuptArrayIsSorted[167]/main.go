package main

func main() {}

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	res := numbers[r] + numbers[l]
	for res != target {
		if res < target {
			l++
		} else {
			r--
		}
		res = numbers[r] + numbers[l]
	}
	return []int{l + 1, r + 1}
}
