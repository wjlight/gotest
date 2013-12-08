package main

import (
	"flag"
	"fmt"
	"math"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list 

	fmt.Println("wawa")
	a, b := MySqrt(9.0)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a, b)

	deferTest()
	bTes()

	fmt.Println("##############")

	bbTest()

	//    if *versionFlag {
	//        fmt.Println("Version:", APP_VERSION)
	//    }
}

func MySqrt(f float64) (v float64, ok bool) {
	if f >= 0 {
		v, ok = math.Sqrt(f), true
	}
	return
}

func deferTest() {
	//go是按先入后出（LIFO）次序，执行一组defer函数
	//你可以在最后关闭所有文件描述符以及解锁所有互斥锁
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

//用defer跟踪代码
func trace(s string)   { fmt.Println("entering", s) }
func untrace(s string) { fmt.Println("leavint", s) }

func aTest() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func bTes() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	aTest()
}

func tra(s string) string { fmt.Println("entering", s); return s }
func unTra(s string)      { fmt.Println("leaving ", s) }

func aaTest() {
	defer unTra(tra("a"))
	fmt.Println("in a")
}

func bbTest() {
	defer unTra(tra("b"))
	fmt.Println("in b")
	aaTest()
}
