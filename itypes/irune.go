package main

import (
	"fmt"
	"unicode/utf8"
)

//  type rune = int32


func testRune() {
	// rune 可以进行整形运算
	r1 := rune('高')
	r2 := rune('a')
	r3 := rune('z')
	fmt.Println(r1)
	fmt.Println(r3 -r2)


	// 字符个数和字符串长度是并不一样的
	str := "hello你好"
	fmt.Println("len:", len(str)) // len 计算的是字节数量，一个汉字占用3个字节
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	fmt.Println(str[1])


}


