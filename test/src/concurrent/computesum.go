package main

import (
	"fmt"
	"runtime"
)

//按cpu进行并行加总执行

type Vector []float64

func (v Vector) Op(f float64) float64 {

}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	//通知完成
	c <- 1
}

const NCPU = 2

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) //接收每个gorutine完成信号

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	//阻塞等待所有完成
	for i := 0; i < NCPU; i++ {
		<-c
	}

	fmt.Println("all finish")
}

func main() {
	runtime.GOMAXPROCS(NCPU) //手动设定使用cpu数量

	runtime.Gosched() //主动出让时间片给其他goroutine
}
