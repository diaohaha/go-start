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

// 回溯算法要素  收集结果 减枝  递归搜索（通常是DFS）
func leetcode22(n int) (res []string) {
	var backtrace func(l int, r int, item string)
	// golang 不支持函数内直接定义函数
	backtrace = func(l int, r int, item string) {

		// 收集结果
		if l == 0 && r == 0 {
			res = append(res, item)
			return
		}

		// DFS
		if l > 0 {
			item += "("
			backtrace(l-1, r, item)
			item = item[:len(item)-1]
		}

		if r > 0 && r > l { // r<l时 剪枝
			item += ")"
			backtrace(l, r-1, item)
		}

		return
	}

	backtrace(n, n, "")
	return
}
