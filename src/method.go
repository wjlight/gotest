package main

import (
	"fmt"
)

func main(){
	testMethor()
	
	test()
	
	testNoNaming()
}


//Go没有类(class)，不过你可以为任何类型附上方法。
//没错，(几乎是)任何类型。
//方法声明为一个带有显式接收者的函数，
//其声明独立于类型的声明
type Point struct{
	x,y int
}
// *Point为显示的接受者 ，
//注意也可以将以Point的形式传入,只是这种值的方式传入，代价比较大
func (p *Point) Alength() int{
	return p.x + p.y
}

func testMethor(){
	p:= Point{2,3}
	fmt.Println(p.Alength())
}

type IntVector [] int
func (v IntVector) Sum()(s int){
	for _,x := range v{
		s += x
	}
	return
}

func test(){
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

//匿名字段的方法
type Point2 struct{x,y int}
func (p *Point2) Lents() int{return p.x+ p.y}
//这个机制提供了一个模拟子类和继承效果的简单方式
type NamePoint struct{
	Point2
	name string
}
//也可以重写
func (p *NamePoint) Lents() int{return p.x * p.y}
func testNoNaming(){
	//如果是不是地址引用的话，貌似运行的比较慢
	n:= &NamePoint{Point2{2,4},"wawa"}
	fmt.Println("n.Lents:", n.Lents())
}