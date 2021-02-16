package main

import "fmt"

//数组切片的数据结构可以抽象为以下3个变量：
//一个指向原生数组的指针；
//数组切片中的元素个数；
//数组切片已分配的存储空间

//基于数组，数组切片添加了一系列管理功能，可以随时动态扩充存放空间，并且可以被随意传递而不会导致所管理的元素被重复复制

func testCreateUseArray() {

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5] //基于数组创建切片,myArray[first:last]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()
}

//arr[first:last]，下标从0开始，first包含，last不包含
//str和arr一样
//只不过len是从1开始
func testIndex() {
	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println("len(arr):", len(arr))
	fmt.Println("arr[0:1]:", arr[0:1])
	fmt.Println("arr[1:2]:", arr[1:2])

	var str string = "abcdefg"
	fmt.Println("len(str):", len(str))
	fmt.Println("str[0:1]:", str[0:1])
	fmt.Println("str[1:2]:", str[1:2])
}

//事实上还会有一个匿名数组被创建出来
func testCreateDirect() {
	mySlice1 := make([]int, 5)       //初始元素5个，初始值0
	mySlice2 := make([]int, 5, 10)   //初始元素5个，初始值0，预留10个元素存储空间
	mySlice3 := []int{1, 2, 3, 4, 5} //创建并初始化

	fmt.Println(mySlice1, mySlice2, mySlice3)
}

func testCreateUseSlice() {

	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]

	fmt.Println("cap(oldSlice):", cap(oldSlice))
	fmt.Println("newSlice:", newSlice)

	oldSlice1 := make([]int, 5, 10)
	newSlice1 := oldSlice1[:7] //范围不能超过oldSlice1的cap
	fmt.Println("newSlice1:", newSlice1)
}

//数组切片多了一个存储能力（ capacity）的概念，即元素个数和分配的空间可以是两个不同的值。
//合理地设置存储能力的值，可以大幅降低数组切片内部重新分配内存和搬送内存块的频率，从而大大提高程序性能
//数组切片会自动处理存储空间不足的问题。如果追加的内容长度超过当前已分配的存储空间
//（即cap()调用返回的信息），数组切片会自动分配一块足够大的内存
func testCap() {

	mySlice := make([]int, 5, 10)
	fmt.Println("len:", len(mySlice)) //切片当前所存储的元素个数
	fmt.Println("cap:", cap(mySlice)) //切片分配的空间大小

	mySlice = append(mySlice, 1, 2, 3) //追加
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...) //追加切片，加省略号相当于把mySlice2包含的所有元素打散后传入
}

func testSliceTravel() {

	mySlice := []int{1, 2, 3, 4}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice[", i, "] = ", mySlice[i])
	}

	fmt.Println()

	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] = ", v)
	}
}

//按照较小的进行拷贝
func testCopy() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) //从slice1->slice2只拷贝3个
	copy(slice1, slice2) //从slice2->slice1只拷贝3个
}

func main() {
	//testCreateUseArray()
	testIndex()
	//testSliceTravel()
	//testCap()
	//testCreateUseSlice()
}
