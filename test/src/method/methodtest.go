package main

import (
	"errors"
	"fmt"
	"math"
)

//函数构成代码执行的逻辑结构。函数的基本组成为：关键字func、函数名、参数列表、返回值、函数体和返回语句。
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}

	return a + b, nil //多重返回
}

//a，b相邻且参数类型相同，则省略前面的变量类型声明，ret1和ret2一样
func Add1(a, b int) (ret1, ret2 int) {
	return 0, 1
}

//导入函数所在包，小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用。
func testMethodInvoke() {
	math.Abs(1.11)
}

//形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数，本质上是一个数组切片[]type
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

//不定参数传递
func myfunc3(args ...int) {
	//原样传递
	myfunc(args...)
	//传递片段
	myfunc(args[1:]...)
}

//任意类型的不定参数
func testAnyType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is an string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

//返回值被命名之后，它们的值在函数开始的时候被自动初始化为空
func testMultiReturn(a int) (b int, err error) {
	return
}

//函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。
func testAnonyMethod() {

	f := func(x, y int) int {
		return x + y
	}

	q := func(a int, b int) int {
		return a + b
	}(1, 2) //}后直接跟参数列表表示函数调用

	fmt.Println(f, q)
}

//闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，
// 而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）
// 为自由变量提供绑定的计算环境（作用域）

// 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
//数据还要表示代码。函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回

//Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在
func testClosure() {
	var j int = 5
	//a指向闭包函数
	a := func() (func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()

	j *= 2

	a()
}

func main() {
	//myfunc(1, 2, 3)

	//var v1 int = 1
	//var v2 int64 = 234
	//var v3 string = "hello"
	//var v4 float32 = 1.234
	//testAnyType(v1, v2, v3, v4)

	//b, err := testMultiReturn(1)
	//fmt.Println(b, err)

	//testAnonyMethod()

	testClosure()
}
