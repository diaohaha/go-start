package main

import (
	"fmt"
	"io"
)

// golang 核心interface

/* io.Reader接口定义了Read(p []byte) (n int, err error)方法，我们可以使用它从Reader中读取一批数据。

type Reader interface {
	Read(p []byte) (n int, err error)
}

读取规则:
// 当输入流结束时，调用它可能返回 err == EOF 或者 err == nil，并且n >=0, 但是下一次调用肯定返回 n=0, err=io.EOF
// 常常使用这个方法从输入流中批量读取数据，直到输入流读取到头，但是需要注意的时候，我们应该总是先处理读取到的n个字节，然后再处理error。
*/

// 一个自定义的reader实现
type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}


func alpha(r byte) byte  {
	if (r >= 'A' && r <= 'Z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read (p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char !=0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	copy(p, buf)
	return  n,nil
}

func main()  {
	reader := newAlphaReader("Fake User Can Kids")
	p := make([]byte, 4)
	// 读取reader的基本写法
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()

}