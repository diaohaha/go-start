package main

import (
	"fmt"
)

func bubleSort(nums []int) []int {
	for end := len(nums); end > 0; end-- {
		for begin := 1; begin < end; begin++ {
			if nums[begin] < nums[begin-1] {
				nums[begin], nums[begin-1] = nums[begin-1], nums[begin]
			}
		}
	}
	return nums
}

func selectSort(nums []int) []int {
	for end := len(nums) - 1; end > 0; end-- {
		// find 0-end 的最大值
		max := nums[0]
		idx := 0
		for begin := 1; begin <= end; begin++ {
			if nums[begin] > max {
				max = nums[begin]
				idx = begin
			}
		}
		nums[idx], nums[end] = nums[end], nums[idx]
	}
	return nums
}

func heapSort(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	//2.调整堆结构+交换堆顶元素与末尾元素
	for j := len(nums) - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		adjustHeap(nums, 0, j)
	}
	return nums
}

func adjustHeap(nums []int, i int, length int) []int {
	// 堆的调整  只调整非叶子节点  从后往前
	// 调整内容  子节点都比自己小，如果不是，找最大的子节点和自己交换
	tmp := nums[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		//如果左子结点小于右子结点，k指向右子结点
		if k+1 < length && nums[k] < nums[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if nums[k] > tmp {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}

	nums[i] = tmp

	return nums
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func insertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		// swap 1
		nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
		// find swap idx
		j := 0
		for j < i && nums[j] < nums[i] {
			j++
		}
		if j != i {
			tmp := nums[i]
			copy(nums[j+1:i+1], nums[j:i])
			nums[j] = tmp
		}
		// fmt.Println(nums)
	}
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	if len(nums) == 2 {
		return []int{min(nums[0], nums[1]), max(nums[0], nums[1])}
	}
	temp := make([]int, len(nums))
	mergeSortCore(nums, 0, len(nums)-1, temp)
	return nums
}

func mergeSortCore(nums []int, left, right int, temp []int) {
	if left < right {
		mid := (right + left) / 2
		mergeSortCore(nums, left, mid, temp)
		mergeSortCore(nums, mid+1, right, temp)
		fmt.Println("i am merge:", nums[left:mid+1], " and ", nums[mid+1:right+1])
		merge(nums, left, mid, right, temp)
		fmt.Println("after merge:", nums[left:right+1])
	}
}

func merge(nums []int, left, mid, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			temp[t] = nums[i]
			t++
			i++
		} else {
			temp[t] = nums[j]
			t++
			j++
		}
	}

	for i < mid {
		temp[t] = nums[i]
		t++
		i++
	}

	for j < mid {
		temp[t] = nums[j]
		t++
		j++
	}
	copy(nums[left:right+1], temp[:right-left-1])
	return
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
func main() {
	// numList := []int{9, 8, 1, 7, 332, 6, 0, -1, 11, -2}

	// nums
	// bubleSort(numList)
	// selectSort(numList)
	// heapSort(numList)
	// insertSort(numList)
	// mergeSort(numList)
	a := "abc"
	fmt.Println(a[1:])
}
