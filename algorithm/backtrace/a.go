package main

import (
	"fmt"
	"testing"
)

var res [][]int
var count int

func run() {
	res = make([][]int, 0)
	nums := []int{1, 2, 3, 4, 5}
	tmp := []int{}
	backtrace(nums, 0, tmp)
	fmt.Println(res)
	fmt.Println(count)
}

func backtrace(nums []int, start int, temp []int) {
	count += 1
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtrace(nums, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}

func TestA(t *testing.T) {
	run()
	t.Log("A")
}
