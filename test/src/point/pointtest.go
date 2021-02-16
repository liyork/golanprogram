package main

import "fmt"

//*type 表示指针类型
//&变量 ，获取变量的内存地址，即指针，得到的类型就是*type
//*指针变量，解引用，即得到指针变量对应的变量

func testPoint() {
	a := 3
	var p *int = &a //变量a的内存地址
	var pv int = *p //解除引用操作符，返回该指针指向的变量的值

	fmt.Printf("a的值为 %v, a的指针是 %v, p指向的内存地址的值为 %v\n", a, p, pv)

	a = 1 //内存对应的值变了，指针地址没变
	fmt.Printf("a的值为 %v, a的指针是 %v, p指向的内存地址的值为 %v\n", a, p, pv)
}

func testPoint1() {
	a := 3
	b := 4
	p1 := &a //变量a的内存地址
	p2 := &b //变量b的内存地址
	fmt.Printf("a的值为 %v, a的指针值 %v, p1指向的变量的值为 %v\n", a, p1, *p1)
	fmt.Printf("b的值为 %v, b的指针值 %v, p2指向的变量的值为 %v\n", b, p2, *p2)

	fmt.Println(demo(p1, p2))

	fmt.Printf("a的值为 %v, a的指针值 %v, p1指向的变量的值为 %v\n", a, p1, *p1)
	fmt.Printf("b的值为 %v, b的指针值 %v, p2指向的变量的值为 %v\n", b, p2, *p2)
}

//指针传递
func demo(a, b *int) int {
	*a = 5 //解引用
	*b = 6
	return *a * *b
}

func main() {
	//testPoint()
	testPoint1()
}
