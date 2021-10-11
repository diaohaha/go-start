package main

import "fmt"

/*
https://leetcode-cn.com/problems/sliding-window-median/
*/

type SortedSlice struct {
	nums   []int
	length int
}

func (s *SortedSlice) insert(num int) {
	if len(s.nums) == 0 {
		s.nums = append(s.nums, num)
		return
	}
	if len(s.nums) == 1 {
		if s.nums[0] > num {
			s.nums = append([]int{num}, s.nums[0])
		} else {
			s.nums = append(s.nums, num)
		}
		return
	}
	l := 0
	r := len(s.nums) - 1
	n := len(s.nums)
	for l < r {
		m := (l + r) / 2
		if s.nums[m] > num {
			// 向左
			r = m - 1
		} else if s.nums[m] == num {
			l = m
			break
		} else {
			l = m + 1
		}
	}
	if s.nums[l] > num && l != 0 {
		l = l - 1
	}
	// 后移一位
	// 首位赋值
	if l == n-1 {
		s.nums = append(s.nums, num)
	} else {
		if l == 0 && num < s.nums[0] {
			l = -1
		}
		tmp := s.nums[l+1]
		s.nums[l+1] = num
		for i := l + 2; i < n; i++ {
			tmp2 := s.nums[i]
			s.nums[i] = tmp
			tmp = tmp2
		}
		// 末尾赋值(需要扩展slice)
		s.nums = append(s.nums, tmp)
	}
}

func (s *SortedSlice) delete(num int) {
	l := 0
	r := len(s.nums) - 1
	for l < r {
		m := (l + r) / 2
		if s.nums[m] > num {
			// 向左
			r = m - 1
		} else if s.nums[m] == num {
			l = m
			break
		} else {
			l = m + 1
		}
	}
	s.nums = append(s.nums[:l], s.nums[l+1:]...)
}

func (s SortedSlice) computeMedian() float64 {
	if len(s.nums)%2 == 0 {
		return float64(s.nums[(len(s.nums)/2)]+s.nums[(len(s.nums)/2-1)]) / 2
	} else {
		return float64(s.nums[(len(s.nums)-1)/2])
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var result []float64
	s := SortedSlice{
		nums: []int{},
	}
	for _, num := range nums[:k] {
		s.insert(num)
		fmt.Println("now is:", s.nums)
	}
	result = append(result, float64(s.computeMedian()))
	for i := k; i < len(nums); i++ {
		fmt.Println("delete:", nums[i-k])
		s.delete(nums[i-k])
		fmt.Println("add:", nums[i])
		s.insert(nums[i])
		fmt.Println("now is:", s.nums)
		result = append(result, float64(s.computeMedian()))
	}
	return result
}

func runLeetcode480() {
	nums := []int{1, 4, 2, 4}
	res := medianSlidingWindow(nums, 4)
	fmt.Println("res is:", res)
}
