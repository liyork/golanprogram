package main

import "fmt"

func Count1(ch chan int, i int) {
	ch <- 1 //写入数据到channel,这个channel被读取之前，这个操作是阻塞的
	fmt.Println("Counting:", i)
}

func main() {
	//创建包含10个channel的数组
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i], i)
		//fmt.Println("go count1:", i)
	}

	//fmt.Println("len(chs):", len(chs))
	for _, ch := range (chs) {
		<-ch //读取数据，写入之前此操作也是阻塞的
	}

	fmt.Printf("end...")
}
