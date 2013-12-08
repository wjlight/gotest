package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	testLit()
}

func testLit() {
	items := list.New()
	for _, x := range strings.Split("ABCD", "") {
		items.PushFront(x)
	}
	items.PushBack(9)
	for element := items.Front(); element != nil; element = element.Next() {
		switch value := element.Value.(type) {
		case string:
			fmt.Printf("%s ", value)
		case int:
			fmt.Printf("%d ", value)
		}
	}
	fmt.Println() //prints:
}
