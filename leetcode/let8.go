package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {
	filterRunes := []rune{}
	for _, r := range []rune(s) {
		if unicode.IsDigit(r) {
			filterRunes = append(filterRunes, r)
		} else if r == '+' || r == '-' {
			if len(filterRunes) > 0 {
				break
			}
			filterRunes = append(filterRunes, r)
		} else {
			continue
		}
	}
	if len(filterRunes) == 0 {
		return 0
	}
	var res int64
	if filterRunes[0] == '+' {
		filterRunes = filterRunes[1:]
	}
	fmt.Println(string(filterRunes))
	res, _ = strconv.ParseInt(string(filterRunes), 10, 32)
	return int(res)
}

func main() {
	a := myAtoi(" -42")
	fmt.Println(a)
}
