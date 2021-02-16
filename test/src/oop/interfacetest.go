package main

import (
	"fmt"
	"oop/one"
	"oop/two"
)

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

//无需继承，实现方法即可用接口类型
type File struct {
	// ...
}

func (f *File) Read(buf []byte) (n int, err error)                { return }
func (f *File) Write(buf []byte) (n int, err error)               { return }
func (f *File) Seek(off int64, whence int) (pos int64, err error) { return }
func (f *File) Close() error                                      {}

func testUse() {
	var file1 IFile = new(File)
	var file2 IReader = new(File)
	var file3 IWriter = new(File)
	var file4 ICloser = new(File)

	fmt.Println(file1, file2, file3, file4)
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

//将对象实例赋值给接口，要求对象实例实现了接口的所有方法
func testAssign1() {
	var a Integer = 1
	var b LessAdder = &a //可以，go根据函数func (a Integer) Less(b Integer) bool，自动生成了
	// func (a *Integer) Less(b Integer) bool{return (*a).Less(b)}
	//var c LessAdder = a//报错,type does not implement 'LessAdder' as 'Add' method has a pointer receiver
	//无法自动根据func (a *Integer) Add(b Integer)生成func (a Integer) Add(b Integer) {(&a).Add(b)}，
	//因为(&a).Add()改变的只是函数Add的参数a(值传递被复制了)，对外部实际要操作的对象并无影响，这不符合用户的预期。所以，Go语言不会自动为其生成该函数。

	fmt.Println(b)
}

//将一个接口赋值给另一个接口.在Go语言中，只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值
func testAssign2() {
	var file1 two.IStream = new(File)
	var file2 one.ReadWriter = file1
	var file3 two.IStream = file2

	//Writer方法是IStream的子集，可以将IStream的实例赋值给Writer，以为多余的接口可以不用，Writer就暴露了Write
	var file4 Writer = file1

	var file5 Writer = new(File)
	//file1 = file5 //不可以，因为Writer中只有Write方法，而IStream还有Read，用户用了file1的Read就报错了

	fmt.Println(file3, file4, file5)
}

type Writer interface {
	Write(buf []byte) (n int, err error)
}

var vi interface {
}

func testInterfaceSearch() {
	var file1 Writer = new(File)
	//检查file1指向的对象实例是否实现two.IStream接口
	//接口查询是否成功，要在运行期才能够确定
	if file5, ok := file1.(two.IStream); ok {
		file5.Read(nil)
	}

	//类型查询，interface{}表示任意类型
	var v1 interface{} = &vi
	switch v := v1.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	}
}

//io.ReadWriter组合了Read和Write

//任何对象都满足空接口interface{}
func testAny() {
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = &v2
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}
}

func main() {
	testInterfaceSearch()
}
