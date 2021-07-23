package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	// nums merge
	nums := make([]int, l)
	if len(nums1) == 0 {
		nums = nums2
	} else if len(nums2) == 0 {
		nums = nums1
	} else {
		i, j := 0, 0
		for m := 0; m < l; m++ {
			if nums1[i] > nums2[j] {
				if j < len(nums2) {
					nums[m] = nums2[j]
					j += 1
				} else {
					// 只能从另一个表中去
					nums[m] = nums1[i]
					i += 1
				}
			} else {
				if j < len(nums2) {
					nums[m] = nums1[i]
					i += 1
				} else {
					// 只能从另一个表中去
					nums[m] = nums2[j]
					j += 1
				}
			}
		}
	}
	// compute median
	midInex := l / 2
	if l%2 == 0 {
		return float64(nums[midInex] + nums[midInex+1])
	} else {
		return float64(nums[midInex+1])
	}
}
