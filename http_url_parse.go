package main

import (
	"fmt"
	"log"
	"net/url"
)

func main()  {

	// url参数中有换行符是解析要替换
	urlStr := "http://www.baidu.com?param=123"
	//newS := strings.Replace(urlStr, "\n", "\\n", -1)
	i, err := url.Parse(urlStr)
	log.Println(i.Path)
	log.Println(i.RawPath)
	log.Println(i)
	log.Println(err)

	params := url.Values{
		"user": []string{"哈哈哈哈\n哈哈哈"},
	}

	urlParams := params.Encode()
	log.Println(params.Encode())

	params1, err := url.ParseQuery(urlParams)
	fmt.Println(params1.Get("user"))

}
