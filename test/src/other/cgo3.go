package main

// #cgo pkg-config: png cairo
// #include <png.h>
import "C" //依赖非C标准库的第三方库

//对于常规返回了一个值的函数，调用者可以用以下的方式顺便得到错误码：
//n, err := C.atoi("a234")

//在传递数组类型的参数时需要注意，在Go语言中将第一个元素的地址作为整个数组的起始地址传入
//n, err := C.f(&array[0])
