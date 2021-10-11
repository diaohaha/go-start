package main

import (
	"fmt"
	"math"
)

const V = "aaa"
var W = "aaa"

func main()  {

	// golang 的3种取整方式
	x :=1.1
	fmt.Println(V)
	fmt.Println(W)
	fmt.Println(math.Ceil(x))// 2  向上取整
	fmt.Println(math.Floor(x))// 1  向下取整
	fmt.Println(math.Floor(1.3 + 0.5)) // 四舍五入
	fmt.Println(math.Floor(1.6 + 0.5))


}
