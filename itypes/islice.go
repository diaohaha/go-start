package main

import "fmt"

// golang 切片
// 切片作为函数参数时是引用传递 !!!

func intRange(n int) []int {
	i := 0
	res := []int{}
	for i < n {
		res = append(res, i)
		i++
	}
	return res
}

func test1() {
	// ** 切片是一个引用类型, 因此当引用改变其中元素的值时候，其他的所有引用都会改变该值。
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[:4]
	s2 := a[3:7]
	fmt.Println(s1)
	fmt.Println(s2)
	s1[3] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	return
}

func test2() {
	// golang中数组声明
	var a [4]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	fmt.Println(a)

	// golang中的切片声明
	// 切片和数组的区别是切片可以动态扩充容量
	var s1 = make([]byte, 6)
	fmt.Println(s1)
	var s2 = make([]byte, 5, 10)
	fmt.Println(s2)
	fmt.Println(cap(s2), len(s2))
	s2 = append(s2, s1...)
	fmt.Println(s2)
	fmt.Println(cap(s2), len(s2)) // 32

	return
}

func test3() {
	s := intRange(20)
	s1 := s[:] // 引用一个切片
	fmt.Println(s1)
	s2 := s
	fmt.Println(s2)
	s3 := s[0:10:30] // 从切片或数组 引用指定长度和容量的切片  30-3就是新slice的容量 30不能超过当前slice容量 否则panic
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	s[1] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	ss := intRange(30)

	// 向切片拷贝内容，只能拷贝len长度的内容过去
	copy(s3, ss)
	fmt.Println(s3)

	s = s[:0]  // 清空切片
	s = s[:10] // 截取切片 s=s[10:]  s=s[10:20]

	// 下面的拷贝不会实际生效
	a := intRange(10)
	aa := []int{}
	copy(aa, a)
	fmt.Println(aa)
}

func test4() {
	// 切片的 删除 和 插入

	//删除切片元素remove element at index
	s := intRange(10)
	// ...是golang对可变参数的支持
	// func append(slice []Type, elems ...Type) []Type

	// 删除index为5的元素
	s = append(s[:5], s[6:]...)
	fmt.Println(s)
	// 删除index为0的元素
	s = append(s[:0], s[1:]...)
	fmt.Println(s)
	// 删除最后一个元素
	s = s[:len(s)-1]
	fmt.Println(s)

	// 切片引用循环的index也是从0开始
	fmt.Println("-------切片循环的index-------")
	// 切片循环的index
	s1 := intRange(10)
	fmt.Println(s1)
	fmt.Println(s1[:3])
	for index, _ := range s1[4:] {
		fmt.Print(index, " - ")
	}
	fmt.Println()
	fmt.Println("--------------")

	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	ss := intRange(10)
	index := 5
	rear := append([]int{}, ss[index:]...)
	ss = append(ss[0:index], 450)
	ss = append(ss, rear...)
	fmt.Println(ss)
	fmt.Println("len(ss)=", len(ss), "cap(ss)=", cap(ss))
}

func testIslice() {
	fmt.Println("test1:")
	test1()
	fmt.Println("--------------------")
	//test2()
	//test3()
	test4()
}
