package main

import "fmt"

//字符串也是一种基本类型
func testStringDefine() {

	var str string //声明
	str = "hello world"
	ch := str[0] //第一字符

	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	//str[0] = 'x'//字符串不能修改
}

//参考strings包
func testStringCompute() {

	var x string = "hello"
	var y string = "123"
	z := x + y

	fmt.Printf("x+y:%s \n", z)
	fmt.Printf("len(x):%d \n", len(x))
	fmt.Printf("x[0]:%c", x[0])
}

//字节数组方式遍历
func testStringTravel1() {

	str := "Hello,世界"
	n := len(str) //12，utf8编码，一个中文占用3个字节
	fmt.Println("len:", n)
	for i := 0; i < n; i++ {
		ch := str[i] //通过下标获取字符,类型为byte
		fmt.Println(i, ch)
	}
}

//以Unicode字符遍历
func testStringTravel2() {

	str := "Hello,世界"
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}
}

func main() {
	//testStringDefine()
	//testStringCompute()
	//testStringTravel1()
	testStringTravel2()
}
