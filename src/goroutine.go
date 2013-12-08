package main

import (
//	"time"
	"fmt"
)

var s string

func main(){
//	testGoRoutine()
	hello()
}

func f(){
	fmt.Println(s)
}

func hello(){
	s = "hello"
	go f()
}


func testGoRoutine(){
	go IsReady("tea", 6)
	go IsReady("coffee", 2)
	fmt.Println("i'm waiting....")
}

//启动一个goroutine
func IsReady(what string , minutes int){
	fmt.Println(what)
//	timeDuration := time.Duration(minutes* 1e9) 
//	time.Sleep(timeDuration) //nanosecs
	fmt.Println(what,"is ready")	
}


