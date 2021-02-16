package main

import (
	"fmt"
	"io"
	"reflect"
)

type MyReader struct {
	Name string
}

//实现了io.Reader接口方法
func (r MyReader) Read(p []byte) (n int, err error) {
	fmt.Printf("read..")
	return 1, nil
}

//反射得到type和value，Type为io.Reader， Value为MyReader{"a.txt"}
//Type主要表达的是被反射的这个变量本身的类型信息，而Value则为该变量实例本身的信息
func typeAndValue() {
	var reader io.Reader
	reader = &MyReader{"a.txt"}

	fmt.Println(reader)
}

func testBaseUsage() {
	var x float64 = 3.4
	typeOf := reflect.TypeOf(x)
	fmt.Println("typeOf:", typeOf)
	fmt.Println("typeOf.Kind:", typeOf.Kind())

	valueOf := reflect.ValueOf(x)
	fmt.Println("valueOf.Type:", valueOf.Type())
	fmt.Println("kind is float64:", valueOf.Kind() == reflect.Float64)
	fmt.Println("values:", valueOf.Float()) //得到值
}

//go中所有类型都是值类型，变量在传递给函数时将发生一次复制
func testSetValue() {
	var x float64 = 3.4

	//v1 := reflect.ValueOf(x)//调用ValueOf时，x产生一个副本，ValueOf内部操作的是x的副本，若下面方法允许，
	// 则被修改的是副本，产生困惑，为什么x本身值没变，为此引入可设属性概念(Settability)
	//v1.Set(4.1)

	p := reflect.ValueOf(&x) //使用变量x的指针
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type T struct {
	A int
	B string
}

func testReflectStruct() {
	t := T{203, "hm203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type() //变量名
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i) //变量
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func main() {
	//testBaseUsage()
	//testSetValue()
	testReflectStruct()
}
