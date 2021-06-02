package main

import (
	"fmt"
	"time"
)

func main()  {

	now := time.Now()  // 2012-10-31 15:50:13.793654 +0000 UTC
	nowUnix := time.Now().Unix()

	// 时间转字符串
	nowStr := now.Format("2006-01-02 15:04-05")
	fmt.Printf("now:", nowStr)

	// 时间戳转时间
	now = time.Unix(nowUnix, 0)

	//  字符串转时间 (涉及到时区)
	timeStr := "2020-01-01 10:10:10"
	loc, _ := time.LoadLocation("Local")  // 获取时区
	t, _ := time.ParseInLocation("2006-01-02 15:04-05", timeStr, loc)
	fmt.Println("%v", t)
	
	// 0值比较
	nilTime := time.Time{} //赋零值

}
