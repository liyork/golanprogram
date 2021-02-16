package main

import "fmt"

//数组就是指一系列同一类型数据的集合，每个数据为元素element

//数组长度在定义后就不可更改
func testArrayDefine() {

	var a1 [32]byte                 //长度为32，每个元素为一个字节
	var a2 [2] struct{ x, y int32 } //复杂类型数组
	var a3 [1000]*float64           //指针数组
	var a4 [3][5]int                //二维数组
	var a5 [2][2][2]float64         //等同于[2]([2]([2]float64))

	arrLen := len(a1) //数组的长度是该数组类型的一个内置常量，len()来获取

	fmt.Println(a1, a2, a3, a4, a5, arrLen)
}

func testArrayElementAccess() {

	array := [3]int{1, 2, 3}
	for i := 0; i < len(array); i++ {
		fmt.Println("Element", i, "of array is ", array[i])
	}

	for i, e := range array {
		fmt.Println("Array element[", i, "]=", e)
	}
}

//Go语言中数组是一个值类型（ value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values:", array)
}

func main() {
	//testArrayElementAccess()

	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(), array values:", array)
}
