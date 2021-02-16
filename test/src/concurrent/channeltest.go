package main

import (
	"fmt"
	"time"
)

func channelDefine() {
	//声明：var chanName chan ElementType
	var ch chan int
	var m map[string]chan bool

	fmt.Println(ch, m)
}

func channelCreate() {
	//声明并初始化
	ch := make(chan int)
	fmt.Println(ch)
}

func channelOpt() {
	ch := make(chan int)
	ch <- 1       //写入，阻塞，直到有其他goroutine从这个channel中读取数据
	value := <-ch //从channel中读取，若没有人写入，那么读也是阻塞的直到channel被写入数据
	fmt.Println(value)
}

func selectSample() {
	var chan1 chan int
	var chan2 chan int

	//每个case语句里必须是一个IO操作
	select {
	case <-chan1: //chan1成功读取到数据
		fmt.Println("chan1 read")
	case chan2 <- 1: //成功向chan2写入数据
		fmt.Println("chan2 write")
	default:
		fmt.Println("nothing")
	}
}

//随机向ch中写入0或1
func testSelect() {
	ch := make(chan int, 1)
	for {
		select { //可能两个case不一定谁能成功
		case ch <- 2: //看来和0没关系
			fmt.Println("put 2")
		case ch <- 1:
			fmt.Println("put 1")
		default:
			fmt.Println("default")
		}
		i := <-ch
		fmt.Println("Value received:", i)
		time.Sleep(1 * time.Second)
	}
}

func testChanCache() {
	//创建带缓冲的channel，缓冲区被填完之前不会阻塞
	c := make(chan int, 1024)

	//读取
	for i := range c {
		fmt.Println("Received:", i)
	}
}

func testChanTimeout() {

	var ch chan int
	//i := <-ch
	//要小心上面用法。若没有人写入ch，那么上面读取动作永远无法从ch中读取到数据，导致整个goroutine永远阻塞

	//利用select机制实现channel操作的超时
	timeout := make(chan bool, 1)
	go func() { //实现并执行一个匿名的超时等待函数
		time.Sleep(1e9) //1s
		timeout <- true
	}()

	select {
	case <-ch: //从ch中读取数据
		fmt.Println("read data from ch:")
	case <-timeout: //一直没从ch中读取到数据，但从timeout中读取到了数据
		fmt.Println("timeout")
	}
}

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int //假设管道中传递的数据只是一个整形数
}

//流水式处理数据
//从queue中读取PipeData然后调用data.handler(data.value)得到结果放入data.next的chan中
func handle(queue chan *PipeData) {
	for data := range queue {
		handleValue := data.handler(data.value)
		fmt.Println("handleValue:", handleValue)
		data.next <- handleValue
		//return
	}
}

func testChanPass() {

	pdChan := make(chan *PipeData, 3)

	ch1 := make(chan int, 1)

	pd1 := &PipeData{1, func(a int) int {
		fmt.Println("pd 1", )
		return 1
	}, ch1}
	//time.Sleep(1e9)
	pdChan <- pd1

	ch2 := make(chan int, 1)
	pd2 := &PipeData{2, func(a int) int {
		fmt.Println("pd 2", )
		return 1
	}, ch2}
	//time.Sleep(1e9)
	pdChan <- pd2

	handle(pdChan)

	v1 := <-ch1
	fmt.Println("v1:", v1)

	v2 := <-ch2
	fmt.Println("v2:", v2)
}

func testChanOneway() {
	var ch1 chan int
	var ch2 chan<- float64 //单向的，只能写float64数据
	var ch3 <-chan int     //单向，只能从channel中读取int数据

	ch4 := make(chan int)
	ch5 := <-chan int(ch4) //ch5就是单向读取channel
	ch6 := chan<- int(ch4) //ch6是单向写入的channel

	fmt.Println(ch1, ch2, ch3, ch4, ch5, ch6)
}

//只能对channel进行读，最小权限原则
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

func testChanClose() {
	ch := make(chan int)
	close(ch)

	//从channel中读取时，判断channel是否已经被关闭
	x, ok := <-ch
	if ok {
		fmt.Println("read value:", x)
	} else {
		fmt.Println("channel closed")
	}

}

func main() {
	//testSelect()
	//testChanTimeout()
	testChanPass()
}
