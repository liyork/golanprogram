package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		//go关键字，这次调用会在一个新的goroutine中并发执行，被调用函数返回时，这个goroutine自动结束，若有返回值则被丢弃
		go Add(i, i)
	}
}
