package main

import (
	"fmt"
	"time"
)

//TODO/*
// xxx.Add(time.Hour)一小时后
// xxx.Sub(xxx)两个时间之间的差值
// xxx.Add(xxx)
// xxx.Add(xxx)
//
//
//
// */

func main() {
	//timex()
	startTime()
	//now := time.Now()
	//fmt.Println(now)
	//later := now.Add(time.Hour)
	//fmt.Println(later)

}
func timex() {
	start := time.Now()
	time.Parse()
	now := time.Now()
	min := now.Add(time.Minute) //一小时以后
	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Println(i)
			break
		}
		fmt.Print(i)
	}
	a4 := now.Before(min) //判断时间真假如果now在end之前返回真否则假
	a5 := now.After(min)  //判断时间真假如果now在end之后返回假否则真
	a3 := start.Equal(now) //判断时间是否相同
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a3)//判断时间是否相同
	//fmt.Println(a1)
	end := time.Now()
	a2 := end.Sub(start)    //时间差
	fmt.Println(a2)//两个时间的差
}

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d\n", year, month, day, hour, minute, second)

}
func timeUnix() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current times tamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

//func (t Time) Add(d Duration) Time

func timestampDemo() {
	now := time.Now()
	timestamp := now.Unix()
	//timestamp2 := now.UnixNano()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d\n", year, month, day, hour, minute, second)

}

func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func startTime() { //时间程序，格式化时间，测试程序运行时间
	start := time.Now()
	now := time.Now().Format("2006/01/02 15:04:05")
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			if j == 4 {
				fmt.Println(i, j)
				break
			}
			fmt.Println(i, j)
		}

	}
	fmt.Println(now)
	end := time.Now()
	last := end.Sub(start)
	fmt.Println(last)
}

func formatDemo() { //时间的格式化
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 mon Jan"))
	fmt.Println(now.Format("2006-01-02 04:04.000 PM Mon Jan"))
	fmt.Println(now.Format("2005/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/02/04"))

}
