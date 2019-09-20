package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr    *cronexpr.Expression
		err     error
		nowtime time.Time
		nextime time.Time
	)
	//每5分钟执行一次*/5 * * * *
	/**
		Field name     Mandatory?   Allowed values    Allowed special characters
		----------     ----------   --------------    --------------------------
		Seconds        No           0-59              * / , -
		Minutes        Yes          0-59              * / , -
		Hours          Yes          0-23              * / , -
		Day of month   Yes          1-31              * / , - L W
		Month          Yes          1-12 or JAN-DEC   * / , -
		Day of week    Yes          0-6 or SUN-SAT    * / , - L #
		Year           No           1970–2099         * / , -

		Parse()里面的表达式，有7个*(用逗号隔开7个字符)分别对应上面的秒、分、小时、天、月、周、年
	*/
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	nowtime = time.Now()
	//根据当前时间来计算出下次运行的时间
	nextime = expr.Next(nowtime)
	fmt.Println("nowtime:",nowtime)
	fmt.Println("nexttime:",nextime)
	//等待这个定时器超时
	fmt.Println("subtime :",nextime.Sub(nowtime))
	time.AfterFunc(nextime.Sub(nowtime), func() {
		fmt.Println("被调用了....", nextime)
	})
	time.Sleep(5 * time.Second)
}
