package main

import "io/ioutil"

func main()  {

	// 文件写入

	// 单次全量写入
	content := []byte("测试1\n测试2\n")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}