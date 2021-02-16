package main

import (
	"fmt"
	"log"
)

//你可以给任意类型（包括内置类型，但不包括指针类型）添加(扩展)相应的方法
func testTypeAddMethod() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}

//定义新类型Integer，和int一样，增加了Less方法
type Integer int

//利用go的面向对象特性
func (a Integer) Less(b Integer) bool {
	return a < b
}

//面向过程方式实现
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

//在Go语言中没有隐藏的this指针”这句话的含义是：
//方法施加的目标（也就是“对象”）显式传递，没有被隐藏起来；
//方法施加的目标（也就是“对象”）不需要非得是指针，也不用非得叫this

//Go语言中的面向对象最为直观，也无需支付额外的成本。
// 如果要求对象必须以指针传递，这有时会是个额外成本，因为对象有时很小（比如4字节），用指针传递并不划算
// 只有在你需要修改对象的时候，才必须用指针。由于Add()方法需要修改对象的值，所以需要用指针引用。
func (a *Integer) Add(b Integer) {
	*a += b
}

//调用后得到a的值还是1，因为类型都是基于值传递的。要想修改变量的值，只能传递指针
func (a Integer) Add1(b Integer) {
	a += b
}

//查看$GOROOT/src/net/http/header.go中的Header
//Header类型其实就是一个map，但通过为map起一个Header别名并增加了一系列方法，它就
//变成了一个全新的类型，但这个新类型又完全拥有map的功能

//值语义和引用语义的差别在于赋值，如果b的修改不会影响a的值，那么此类型属于值类型。如果会影响a的值，那么此类型是引用类型。
//Go语言中的大多数类型都基于值语义，包括：
// 基本类型，如byte、 int、 bool、 float32、 float64和string等；
// 复合类型，如数组（ array）、结构体（ struct）和指针（ pointer）等

//Go语言中的数组和基本类型没有区别，是很纯粹的值类型
func testArrayValueType() {
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b)
}

func testPointType() {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println(a, *b)
}

//数组切片：指向数组（ array）的一个区间。数组切片类型本身的赋值仍然是值语义
func testSliceType() {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println(a, *b)
}

//map本质上是一个字典指针
//channel和map类似，本质上是一个指针。将它们设计为引用类型而不是统一的值类型的原因是，完整复制一个channel或map并不是常规需求
//接口具备引用语义，是因为内部维持了两个指针，data *void 和 itab *Itab
//结构体类似于class，go放弃了包括继承在内的大量面向对象特性，只保留组合特性
//所有的Go语言类型（指针类型除外）都可以有自己的方法

//结构体只是普通的符合类型，定义矩形类型
type Rect struct {
	x, y          float64
	width, height float64
}

//成员方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

//go中未进行显式初始化的变量都会被初始化为该类型的零值
func testCreateStruct() {
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
}

//在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名，表示“构造函数”：
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//Go语言也提供了继承，但是采用了组合的文法，所以我们将其称为匿名组合
type Base struct {
	Name string
}

type ABase struct {
	Name string
}

//成员方法
func (base *Base) Foo() {}
func (base *Base) Bar() {}

type Foo struct {
	Base   //清晰展示内存布局
	*ABase //可以以指针从一个类型派生
}

//重写Bar方法，
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func testInherit() {
	var foo = &Foo{}
	foo.Foo() //从组合Base继承的方法
}

type Job struct {
	Command     string
	*log.Logger //组合了log.Logger指针
}

func inheritLog() {
	//在Job类型的所有成员方法中可以很舒适地借用所有log.Logger提供的方法
	var job *Job
	job.Println("aaaa")
}

//同样命名的变量，以最外面为准
//匿名组合类型相当于以其类型名称（去掉包名部分）作为成员变量的名字

//Go语言中符号的可访问性是包一级的而不是类型一级的，即小写只能包内可见。

func main() {
	//testTypeAddMethod()

	//var a Integer = 1
	//if Integer_Less(a, 2) {
	//	fmt.Println(a, "Less 2")
	//}

	//var a Integer = 1
	//a.Add(2)
	//fmt.Println("a = ", a)

	//testArrayValueType()
	testPointType()
}
