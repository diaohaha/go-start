package main

import (
	"fmt"
	"math"
)

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
