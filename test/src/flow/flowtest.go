package main

import "fmt"

func testIf() int {
	a := 1

	//条件不需要()
	if a < 5 { //{必须,且在if所在行
		return 0
	} else {
		return 1
	}
}

//在有返回值的函数中，不允许将“最终的” return语句包含在if...else...结构中，
//否则会编译失败：function ends without a return statement。
//失败的原因在于， Go编译器无法找到终止该函数的return语句。编译失败的案例如下
func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
} //这个似乎没有失败

func testSwitch() {

	var a int = 1
	switch a { //{与switch同行
	case 0:
		fmt.Printf("0") //不需要break
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough //才会继续执行下一个case
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4,5,6")
	default:
		fmt.Printf("Default")
	}

	var Num int
	switch { //不设定switch之后的条件表达式
	case 0 <= Num && Num <= 3:
		fmt.Printf("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Printf("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Printf("7-9")
	}
}

func testFor() {

	sum := 0
	for i := 0; i < 10; i++ { //不用()，{与for同行
		sum += i
	}

	//无限循环
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	//条件表达式中多重赋值，平行赋值初始化
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println("JLoop:", i)
		}
	}

}

func testGoto() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func main() {
	//i := testIf()
	//fmt.Println("i:", i)

	//testFor()
	testGoto()
}
