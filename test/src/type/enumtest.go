package main

//枚举指一系列相关的常量

//大写字母开头的常量在包外可见
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays //包内私有，常量没有导出
)
