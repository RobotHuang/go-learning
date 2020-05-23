package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\t%v\n", now, now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1, timestamp2)
	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)

	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}

	//格式必须使用这个时间2006年1月2日15点05分0秒 周一 一月而不是yyyy-dd-mm HH:MM:SS
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj1, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	fmt.Println(timeObj1)

	end := time.Now()
	duration := end.Sub(now)
	fmt.Println(duration)
}