package main //表示所属包

import (
	"fmt" //依赖，必须有用
)

//要生成go可执行程序，必须建立一个名字为main的包，并且该包中包含一个叫main()的函数，main函数是go可执行程序的执行起点
func main() {
	fmt.Printf("hello world")
}

//函数定义
//func 函数名(参数列表)(返回值列表){
//	//函数体
//}

/*多返回值，未赋值的则被设置为默认值result=0.0,err=nil*/
func Compute(value1 int, value2 float64) (result float64, err error) { //必须放这里
	//函数体
	return //go并不要求每个语句后面加上分号
}

//cd ~/goyard
//go run hello.go  #直接运行(编译、链接、运行)

//go build hello.go  #编译
//./hello  #运行

//go命令行工具只是一个源代码管理工具，一个前端。真正的go编译器和链接器在后面
//6g hello.go  #64位go编译器
//61 hello.6  #64位go链接器
//./6.out
