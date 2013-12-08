package main

import (
	"fmt"
)

func main() {
	//	test()
	// pumpTest()
	// testPump2()
	chanDirTest()
}

//一个channel创建时，都是双向的，但是我们可以赋值给带方向性的channel变量
//
func channnelDir() {
	var recvOnly <-chan int
	var sendOnly chan<- int
}
func sink(ch <-chan int) {
	for {
		<-ch
	}
}
func source(ch chan<- int) {
	for {
		ch <- 1
	}
}
func chanDirTest() {
	c := make(chan int) //双向
	go source(c)
	go sink(c)
}

//返回channel的函数，是Go中一个重要的惯用法
func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func testPump2() {
	stream := pump2()
	fmt.Println(<-stream) // 打印0
}

func test() {
	c := make(chan int)
	c <- 1 //向c发送1
	fmt.Println("c:", c)
	var v = <-c //从c中，接受数据，赋值给v
	fmt.Println("v:", v)
	<-c      //接受数据，丢弃
	i := <-c //接受值，用于初始化i
	fmt.Println("i:", i)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func pumpTest() {
	ch1 := make(chan int)
	go pump(ch1) //pump挂起，我们运行
	fmt.Println(<-ch1)
	go suck(ch1)
}

//启动一个循环接受者
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
