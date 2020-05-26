// 时间日期相关的函数 time包
package main

import (
	"fmt"
	"time"
)

func main() {

	// 获取当前的时间
	now := time.Now()
	fmt.Printf("now=%v nowtype=%T \n", now, now) //now=2019-03-28 21:48:32.1547 +0800 CST m=+0.000202093 nowtype=time.Time

	// 获取年月日信息
	fmt.Println("year: ", now.Year())
	fmt.Println("month: ", int(now.Month()))
	//...
	fmt.Println("second: ", now.Second())

	//格式化日期
	//1:
	dateFormat := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateFormat)
	//2: 这个时间必须是这个日期，否则会有问题，这个日期是go作者想起来做go的时间点。
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("01")) //只取月

	// 时间的常量
	/*
	const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	 */
	for i := 1; i < 3; i++ {
		fmt.Println(i)
		// 每秒钟输出，需要使用时间的常量
		time.Sleep(time.Second)
	}

	//unix时间戳 nano时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

}
