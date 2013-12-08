package main

import (
	"fmt"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filer(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		fmt.Println("i:", i, " prime:", prime)
		if i%prime != 0 {
			out <- i
		}
	}
}

//筛子
func main() {
	ch := make(chan int)

	go Generate(ch)

	for i := 0; i < 3; i++ {
		prime := <-ch
		fmt.Println("prime:", prime)

		out := make(chan int)
		go Filer(ch, out, prime)
		ch = out
	}
}
