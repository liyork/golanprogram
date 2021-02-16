package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

//漂亮的错误处理规范是Go语言最大的亮点之一。

//返回错误
func Foo(param int) (n int, err error) {
	return
}

func fooInvoke() {

	n, err := Foo(0)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("right n:", n)
	}
}

//自定义error类型，error.go中

//承载错误信息的类型
type PathError struct {
	Op   string
	Path string
	Err  error
}

//实现Error()方法，编译器知道PathError可以当一个error传递
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

//使用PathError作为error
func Stat(name string) (fi int, err error) {
	var stat syscall.Stat_t

	err = syscall.Stat(name, &stat)

	if err != nil {
		return 1, &PathError{"stat", name, err}
	}

	return 1, nil
}

//defer
//即使有异常则也能正常srcFile/dstFile的Close
//defer执行是后进先出
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	defer func() {
		//clean multi statement
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//panic
//当在一个函数调用panic()函数时，执行defer然后退出当前函数，并导致逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止。
//错误信息将被报告，包括在调用panic()函数时传入的参数，这个过程称为错误处理流程。
func testPanic() {
	fmt.Println("testPanic start")
	testPanic1()
	fmt.Println("testPanic end") //不会执行
}

func testPanic1() {
	fmt.Println("testPanic1 start")
	panic("test panic")
	fmt.Println("testPanic1 end") //不会执行
}

//recover()函数用于终止错误处理流程。一般情况下， recover()应该在一个使用defer
//关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
//过程（使用recover关键字），会导致该goroutine所属的进程打印异常信息后直接退出。
func testRecover() {
	//担心testPanic()有错误，defer肯定会执行
	defer func() {
		if r := recover(); r != nil { //recover使得错误处理过程终止
			fmt.Printf("Runtime error caught: %v", r)
		}
	}()

	testPanic()
	fmt.Printf("testRecover end") //不执行
}

func main() {
	//_, err := Stat("a.txt")
	////获取错误相信信息，类型转换
	//if err != nil {
	//	if e, ok := err.(*os.PathError); ok && e.Err != nil {
	//		//获取PathError类型变量e中的其他信息并处理
	//	}
	//}

	//testPanic()
	testRecover()
}
