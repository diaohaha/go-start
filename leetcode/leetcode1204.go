package main

import (
	"fmt"
	"math"
)

/*
question: https://leetcode-cn.com/problems/get-equal-substrings-within-budget/
当框内总消耗小于消耗阈值时，延展右侧窗口。
当框内纵消耗大于消耗阈值时，当前长度定为最大值；右侧移动一格，左侧也随之移动一格，维持窗口为最大符合条件的大小。
直到可以继续扩充时，才扩充右侧大小，所以在整个滑动过程中，窗口的大小是依据合法性只增不减，且直到数据末尾时的窗口大小为所有子串中的最大合法窗口，且右指针与n重合，窗口大小为n-l。
使用窗口合法非递减的性质，利用左右指针保证了算法的正确性。妙绝妙绝。
*/

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	l, r := 0, 0
	usedCost := float64(0)
	for (r < n) {
		usedCost += math.Abs(float64(rune(t[r]) - rune(s[r])))
		r += 1
		if (usedCost > float64(maxCost)) {
			usedCost -= math.Abs(float64(rune(t[l]) - rune(s[l])))
			l += 1
		}
	}
	return n-l
}

func printInt(s string) {
	for _, i := range s {
		fmt.Print(rune(i), " ")
	}
	fmt.Println()
}

func main()  {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	printInt("pxezla")
	printInt("loewbi")
	fmt.Println(equalSubstring("pxezla", "loewbi", 25))
}
