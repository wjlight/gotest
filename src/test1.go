package main

import "fmt"

func chanledTest() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println("hello world2")
	// go chanledTest()
	// dealString()
	rangeTest()
}

//go中的string是不可以修改的，要修改可以通过如下
func dealString() {
	s := "hello world"
	//转成rune数组, rune是int32的别名
	c := []rune(s)
	c[0] = 'w'
	s2 := string(c)
	fmt.Println(s2)
}

//range，打散成独立的unicode字符，按照UTF-8解析
func rangeTest() {
	s := "waw哇哇"
	for k, v := range s {
		fmt.Printf("cha '%c' starts at byte position %d\n", v, k)
	}
}

func switchTest() {
	ca = 0
	switch ca {
	case 0:
		fallthrough //当ca 为 0 的时候，如果没有fallthrough，go是不会自动向下走到case1的地方的
	case 1:
		fmt.Println("gaga")

	case '', '?','+':
		fmt.Println("多个")
	default:
		fmt.Println("default")
	}
}
