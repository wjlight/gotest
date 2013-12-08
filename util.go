package main 

import (
	"fmt"
)

func main() {
	var arrays [3] int
	isZero := check(0)
	fmt.Println("",isZero)
	
	aa1 := [3]int{1,3,4}
	fmt.Println("aa1:",aa1)
	
	aa2 := [...]int{1,2,3}
	fmt.Println("aa2:",aa2)
	
	aa3 := [10]int{2:1, 3:3, 6:2}
	fmt.Println("aa3:",aa3)
	
	sliceTes()
	
	
	lenS := len(arrays)
	fmt.Println("len:",lenS)
	
	mapTest()
	
	ergodicMap()
	
	structTest()
	
	testMember()
}
/*
new是为了分配内存。
注意make([]int, 10)返回[]int，
而new([]int)返回*[]int
*/

type Point struct{	
	x,y int
}

func structTest(){
	var p Point
	p.x = 7
	p.y = 5
	fmt.Println("p:", p)  //7,5
	var pp *Point = new(Point)
	//struct是值类型
	*pp = p
	//结构体指针，没有->符号可用
	pp.x = 9   //(*pp).x的语法糖
	fmt.Println("p2:", p) //7,5
	fmt.Println("pp:", pp) // 9,5
	fmt.Println("pp:", &pp)
	
	p = Point{3,4}
	fmt.Println("p:", p)
	
	p = Point{x:4,y:8}
	fmt.Println("p2:", p)
	
	pp = &Point{5,6}
	pp = &Point{}//等价于new(Point)
}

type A struct{
	ax,ay int
}
//任意具名类型或指向具名类型的指针都可以用作匿名字段
//他们可以出现在结构体中的任意位置
type C struct{
	x float64
	int
	string
}
type B struct{
	A   //匿名结构体字段
	bx,by int
}

func testMember(){
	//匿名字段,B的字面量必须提供细节
	b := B{A{1,2},3,5}
	fmt.Println("b:",b)
	fmt.Println("b.A:",b.A)
}

func mapTest(){
	m := map[string]int{"name":1,"addr":2}
	fmt.Println("m:",m)
	
	m = make(map[string]int)
	fmt.Println("m2:",m)
	
	one := m["name"]
	fmt.Println("one:",one)
	
	
	var value int
	var has bool
	value, has = m["awawa"]
	fmt.Println("value:",value, has)
	
	m = map[string]int{"name":1,"addr":2}
	delete(m, "name")
	fmt.Println("m3:", m)
	
}

func ergodicMap(){
	m := map[string]int{"name":1,"addr":2}
	for key, value:= range m{
		fmt.Println("key:", key, "value:", value)	
	}
}

func sliceTes(){
	//切片
	var slice = make([]int, 10)
	fmt.Println("slice:",slice)
	
	//长度为0，容量为10的切片,切片的传递仅需要很小的代价
	var slice2 = make([]int,0,10)
	fmt.Println("slice2:",slice2)
	
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[5:6]
	fmt.Println("a 5~6:",a)
	a = a[0:4]
	fmt.Println("a 0~4:",a)
}

func check(isZero int) (bool){
	if isZero == 0 {
	 	return true
	} 
	return false
}

