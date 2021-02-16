package main

func testBool() {
	var v1 bool
	v1 = true
	v2 := (1 == 2) //被推导为bool

	//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	var b bool
	//b = 1       //编译错误
	//b = bool(1) //编译错误
}
