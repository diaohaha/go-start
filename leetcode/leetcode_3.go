package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lengthMap := make(map[int]int, len([]rune(s)))
	for i, r := range []rune(s) {
		// compute lengthMap[i]
		fmt.Println(r)
		length := 0
		tmp := map[rune]bool{}
		for j := i; j < len([]rune(s)); j++ {
			fmt.Println(tmp)
			if _, ok := tmp[[]rune(s)[j]]; ok {
				break
			} else {
				tmp[[]rune(s)[j]] = true
				length += 1
			}
		}
		lengthMap[i] = length
	}
	max := 0
	for _, l := range lengthMap {
		if l > max {
			max = l
		}
	}
	return max
}

func runLeetcode3() {
	i := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(i)
}
