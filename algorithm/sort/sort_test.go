package test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"
)

// insert buble select 是3个入门排序算法
// shell排序 缩小增量排序

func QuickSort(arr []int, left, right int) {
	if !(left < right) {
		return
	}
	i := left
	j := right
	temp := arr[i]
	for i < j {
		for i < j && arr[j] >= temp {
			j -= 1
		}
		if i < j {
			arr[i] = arr[j]
			i += 1
		}

		for i < j && arr[i] < temp {
			i += 1
		}
		if i < j {
			arr[j] = arr[i]
			j -= 1
		}
	}
	arr[i] = temp
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

func genArr(n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(100000))
		arr = append(arr, int(n.Int64()))
	}
	return arr
}

// go test -v algorithm/sort/sort_test.go -test.run TestQuickSort
func TestQuickSort(t *testing.T) {
	arr := genArr(1000000)
	QuickSort(arr, 0, len(arr)-1)
	// t.Log(arr)
}

func BubleSort(arr []int, begin, end int) {
	var isSwap bool
	for j := end; j > begin; j-- {
		isSwap = false
		for i := begin; i < j; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

// go test -v algorithm/sort/sort_test.go -test.run TestBubleSort
func TestBubleSort(t *testing.T) {
	arr := genArr(100000)
	// t.Log(arr)
	BubleSort(arr, 0, len(arr)-1)
	// t.Log(arr)
}

// 插入排序
func InsertSort(arr []int, num int) {
	for j := 1; j < num; j++ {
		// 从第2位开始插入
		candidate := arr[j]
		insertP := 0
		for (insertP < j) && candidate > arr[insertP] {
			insertP += 1
		}

		// 将candidate插入在place
		i := j
		for i > insertP {
			arr[i] = arr[i-1]
			i -= 1
		}
		arr[insertP] = candidate
	}
}

// go test -v algorithm/sort/sort_test.go -test.run TestInsertSort
func TestInsertSort(t *testing.T) {
	arr := genArr(100)
	InsertSort(arr, len(arr))
}

// go test -v algorithm/sort/sort_test.go -test.run TestTimeConsume
func TestTimeConsume(t *testing.T) {
	originArr := genArr(100000)
	var sTime, eTime int64
	sTime = time.Now().UnixMilli()
	arr := make([]int, len(originArr))
	copy(arr, originArr)
	fmt.Println("len:", len(arr))
	InsertSort(arr, len(arr))
	eTime = time.Now().UnixMilli()
	t.Log("insert sort consume:", eTime-sTime, "ms")

	sTime = time.Now().UnixMilli()
	copy(arr, originArr)
	BubleSort(arr, 0, len(arr)-1)
	eTime = time.Now().UnixMilli()
	t.Log("buble sort consume:", eTime-sTime, "ms")
}
