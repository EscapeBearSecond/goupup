package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间
	fmt.Println(now)

	// 时间格式化
	layout := "2006-01-02 15:04:05"
	fmt.Println(now.Format(layout))

	// 字符串转换成时间
	timeString := "2023-03-29 16:25:47"
	t1, err := time.Parse(layout, timeString)
	if err != nil {
		panic(err)
	}
	fmt.Println("t1:", t1)
	// 字符串转换成时间，并指定时区
	local, err := time.LoadLocation("Asia/Shanghai")
	t2, err := time.ParseInLocation(layout, timeString, local)
	if err != nil {
		panic(err)
	}
	fmt.Println("t2:", t2)
	// 时区
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(loc))
	// 时间运算
	fmt.Println(now.Add(1 * time.Hour)) // 加时间
	fmt.Println(now.AddDate(1, 1, 1))   // 加日期
	t3 := now.AddDate(1, 1, 1)
	fmt.Println(now.Sub(t3)) // 两个时间相减

	// 时间的比较
	now.Equal(t1)  // =
	now.After(t1)  // >
	now.Before(t1) // <

	// 时间戳
	fmt.Println(now.Unix())      // 时间戳秒
	fmt.Println(now.UnixMilli()) // 毫秒
}
