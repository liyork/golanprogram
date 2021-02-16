package main

import (
	"fmt"
)

func testIntShow() {
	var value2 int32
	value1 := 64 //推导为int
	//value2 = value1 //编译错误，int和int32是两种不同的类型
	value2 = int32(value1) //通过，需要注意数据精度损失和值溢出问题。
}

//+、-、*、/和%
func testIntCompute() {

	fmt.Println(5 % 3) //取余，2
}

//>、 <、 ==、 >=、 <=和!=
func testIntCompare() {

	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal.")
	}

	var i32 int32
	var i64 int64
	i32, i64 = 1, 2
	//if i32==i64{//编译错误，两个不同类型的整型数不能直接比较
	//}

	//各种类型的整型变量都可以直接与字面常量（ literal）进行比较
	if i32 == 1 || i64 == 2 {
		fmt.Println("i and j are equal.")
	}
}

func testIntBitCompute() {
	//124<<2=496、124>>2=31、124^2=126、124&2=0、124|2=126、^2=-3(取反)
}
