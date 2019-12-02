package main

import (
	"fmt"
	"time"
)

/**
	时区操作
 */
const TIME_LAY_OUT1 = "2006-01-02 15:04:05" //这是固定写法
const TIME_LAY_OUT2 = "2006-01-02 03:04:05 PM"
const TIME_LAY_OUT3 = "02/01/2006 15:04:05 PM"

//根据当前日期获取日期的基本属性
func demo1() {
	t := time.Now()   //得到的是一个time.Time对象
	fmt.Println("当前时间的时间戳：",t.Unix())
	fmt.Println("当前时间的时区：",t.Location())
	fmt.Println("当前时间输出",t.String())
	fmt.Println("当前时间的秒数",t.Second())
	fmt.Println("当前时间的分数",t.Minute())
	fmt.Println("当前时间的小时",t.Hour())
	fmt.Println("当前时间的月份天数",t.Day())
	fmt.Println("当前时间的月份",t.Month())
	fmt.Println("当前时间的年份",t.Year())
	fmt.Println("当前时间的年份天数",t.YearDay())
	fmt.Println("当前时间的星期几",t.Weekday())
	fmt.Println(t.Clock()) //获取当前时间的时、分、秒这三个返回值
	fmt.Println(t.Date())
}

//时间戳转换为日期格式
func demo2() {
	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format(TIME_LAY_OUT1))
	fmt.Println(tm.Format(TIME_LAY_OUT1))
	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())
}

//字符串时间转换为时间戳
func demo3()  {
	//time.LoadLocation("Asia/Shanghai")   //这个在Windows下面执行没有效果，这个依赖于tzdata这个包，在Windows下面没有
	//t1,_ := time.Parse(TIME_LAY_OUT1,"2019-05-23 17:02:50")  //不能使用这个，因为在使用这个的时候，是使用的时区是UTC，要使用下面time.ParseInLocation()这个函数，并指定时区，这样就不会出错。
	t1,_ := time.ParseInLocation(TIME_LAY_OUT1,"2019-05-23 17:02:50",time.Local)
	fmt.Println(t1.Location())
	tm := t1.Unix()
	fmt.Println(tm)
}

//获取指定日期的前后整点值
func demo4()  {
	t := time.Now()
	// 整点（向下取整）就是获取当前时间的最大小时整点时间。这里取决于Truncate()里面的单位
	fmt.Println(t.Truncate(1*time.Hour))
	// 整点（向上取整） 就是如果当前时间分钟超过30，就会获取当前小时的下一个小时的整点时间。
	fmt.Println(t.Round(1*time.Hour))
}

//日期的计算
func demo5()  {
	t := time.Now()
	nt := t.Add(2*time.Hour)
	fmt.Println(nt.Format(TIME_LAY_OUT1))

}

func main() {
	demo1()
}
