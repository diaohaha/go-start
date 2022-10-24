package main

func splitList(nums []int) (res int) {
	ptr := 0
	maxLeftVal := nums[0]

	ptr += 1
	for ptr < len(nums) {
		temp := maxLeftVal
		for j := ptr + 1; j < len(nums); j++ {
			if nums[j] < maxLeftVal {
				ptr = j
				maxLeftVal = temp
				break
			} else {
				if nums[j] > temp {
					temp = maxLeftVal
				}
			}
		}
	}

	return ptr + 1
}
