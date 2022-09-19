package main

import "fmt"

var resultCount int
var n, k, s int

func main() {
	// n k s
	resultCount = 0
	n = 3
	k = 6
	s = 5
	backtrace2(0, 0)
	fmt.Println("result count:", resultCount)
}

func backtrace2(tmpSum, i int) {
	fmt.Println("now get ", i, "numbers,", "sum is:", tmpSum)
	if i > n {
		return
	}
	if tmpSum == s {
		if i == n {
			resultCount += 1
			fmt.Println("run here")
			return
		} else if i < n {
			return
		}
	} else if tmpSum < s {
		if i == n {
			return
		} else if i < n {
		}
	} else {
		// tmpSum > s
		return
	}

	for start := 1; start <= k; start++ {
		tmpSum += start
		backtrace2(tmpSum, i+1)
		tmpSum -= start
	}
}
