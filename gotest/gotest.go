package gotest

import (
	"errors"
	
)

func Division2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}

func testLit() {
	items := list.New()
	for _, x := range strings.Split("ABCD", ""){
		items.PushFront(x)
	}
	items.PushBack(9)
	for element := items.Front(); element != nil{
		element = element.Next(){
			switch value := element.Value.(type){
			case string:
				fmt.Println("%s ", value)
			case int:
				fmt.Println("%d ", value)
			}
		}
	}
	fmt.Println()  //prints:

}
