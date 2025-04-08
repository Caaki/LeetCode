package main

func main() {

}

func minimunOperations(nums []int) int {
	vals := make([]int8, 101)
	var added int8 = 0

	for i := len(nums) - 1; i >= 0; i-- {
		if vals[nums[i]] != 0 {
			break
		} else {
			vals[nums[i]] = int8(nums[i])
			added++
		}
	}

	if added == 0 {
		return 0
	}
	if (int8(len(nums))-added)%int8(3) != 0 {
		return int((int8(len(nums))-added)/3) + 1
	}
	return int((int8(len(nums)) - added) / 3)
}
