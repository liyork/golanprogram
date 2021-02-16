package main

import (
	"fmt"
	"sync"
)

//全局只运行一次
//once的Do()方法可以保证在全局范围内只调用指定的函数一次（这里指setup()函数），而且所有其他goroutine在调用到此语句时，
// 将会先被阻塞，直至全局唯一的once.Do()调用结束后才继续

var a string
var once sync.Once

func setup() {
	a = "Hello, world"
	fmt.Println("setup exec")
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

//sync/atomic中有相关原子操作，如CompareAndSwapInt64

func main() {
	fmt.Println("main exec")
	twoprint()
}
