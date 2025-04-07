package main

func main() {

}

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	values := make(map[int]bool)
	values[0] = true

	for _, v := range nums {
		if v == target {
			return true
		}
		nVals := make(map[int]bool, 0)
		for k, _ := range values {
			if k+v == target {
				return true
			}
			nVals[k+v] = true
			nVals[k] = true
		}
		values = nVals
	}
	return false
}
