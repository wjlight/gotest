package main

import (
	"fmt"
	"time"
)

//定时器，跟sleep类似，但是它可以stop
func main() {
	// timerTest()
	sampleEg()
}

func timerTest() {
	timer := time.NewTimer(2 * time.Second)

	//只能通知一次啊
	<-timer.C
	fmt.Println("time expired!")
}

//使用Ticker，可以多次通知，下面的代码是一个死循环
func mulNotified() {
	ticker := time.NewTicker(time.Second)

	//多次通知
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}

func sampleEg() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("wawa", t)
		}
	}()

	//10秒后，停掉Ticker
	timer := time.NewTimer(10 * time.Second)

	<-timer.C

	ticker.Stop()

	fmt.Println("timer expired!")
}
