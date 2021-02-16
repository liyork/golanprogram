package main

//变量相当于是对一块数据存储空间的命名，程序可以通过定义一个变量来申请一块数据存储空间，之后可以通过引用变量名来使用这块存储空间。

func varDefine() {
	//纯粹的变量声明，使用var，类型信息放在变量名之后，语句不用分号结束
	var v1 int
	var v2 string
	var v3 [10]int //数组
	var v4 []int   //数组切片
	var v5 struct {
		f int
	}
	var v6 *int           //指针
	var v7 map[string]int //map，key为string，value为ini
	var v8 func(a int) int

	//一起定义
	var (
		v9  int
		v10 string
	)
}

func varInit() {

	//声明变量时直接初始化var可以去掉
	var v1 int = 10
	var v2 = 10 //编译器自动推导
	v3 := 10    //":="符号，明确表达同时进行变量的声明和初始化，编译器从右边推导出变量类型，之前不允许被声明过
}

func varAssign() {

	var v10 int
	v10 = 1234

	//多重赋值
	var j = 1
	var i = 2
	i, j = j, i //不需要额外引入中间变量
}

func varAnonymity() {

	//只想获得nickName
	_, _, nickName := getName()
}

//定义三个返回值
func getName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}
