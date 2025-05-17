package main

func main() {}

func sortColors(nums []int) {
	zeroPos := 0
	twoPos := len(nums) - 1
	if len(nums) == 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && i == zeroPos {
			zeroPos++
			continue
		}
		if nums[i] == 0 {
			for zeroPos < len(nums)-1 && nums[zeroPos] == 0 {
				zeroPos++
			}
			if zeroPos >= len(nums)-1 {
				break
			}
			if zeroPos-1 == i {
				continue
			}
			nums[i], nums[zeroPos] = nums[zeroPos], nums[i]
			zeroPos++

			continue
		}

		if nums[i] == 2 {
			for twoPos >= 0 && nums[twoPos] == 2 {
				twoPos--
			}
			if twoPos <= i {
				break
			}

			nums[i], nums[twoPos] = nums[twoPos], nums[i]
			twoPos--
			if nums[i] == 0 {
				i--
			}
		}
	}
}
