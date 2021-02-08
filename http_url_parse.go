package main

import (
	"net/url"
	"strings"
)

func main()  {

	// url参数中有换行符是解析要替换
	urlStr := "/service?user_id=&vote_id=&user=abc\n和孙悟空&&source=1"
	newS := strings.Replace(urlStr, "\n", "\\n", -1)
	_, _ = url.Parse(newS)
}
