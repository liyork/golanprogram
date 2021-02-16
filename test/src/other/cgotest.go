package main

import "C"
import (
	"fmt"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C" //信号，让Cgo工作，对上面块注释中(一定要紧贴import)的C源代码自动生成包装性质的go代码，Cgo生成的代码是帮你封装了压栈和调用的过程

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func testInvokeC() {
	Seed(100)
	fmt.Println("Random:", Random())
}

//go的string对应c的字符数组。Cgo提供一些列函数支持：C.CString、C.GoString和C.GoStringN，每次转换都导致一次内存复制
//C.CString的内存管理方式与go语言自身的管理方式不兼容，所以在使用完后必须显示释放调用C.CString所生成的内存块
func testFree() {
	var gostr string
	gostr = "abc"
	cstr := C.CString(gostr)
	defer C.free(unsafe.Pointer(cstr))

	//C.sprintf(cstr, "content is: %d", 123)
}

func main() {
	//testInvokeC()
	//testFree()
}

//go run cgotest.go
