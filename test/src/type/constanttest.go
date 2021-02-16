package main

//常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等。

//字面常量,程序中硬编码的常
func literal() {

	//-12       //无类型，只要常量在相应类型的值域范围内，就可作为该类型的常量
	//3.14159   //浮点类型的常量
	//3.2 + 12i //复数类型的常量
	//true      //布尔常量
	//"foo"     //字符串常量
}

func constantDefine() {

	const PI float64 = 3.14159
	const zero = 0.0 //无类型浮点常量
	const (
		size int64 = 1024
		eof        = -1 //无类型整型常量
	)
	const u, v float32 = 0, 3   //u=0.0, v=3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo" //a=3,b=3,c="foo"，无类型整形和字符串常量

	const mask = 1 << 3 //编译期运算的常量表达式

	//const home = os.Getenv("HOME")//报错，常量赋值是一个编译期行为，右值不能出现任何需要运行期才能得到的结果表达式
}

//预定义了这些常量： true、 false和iota
func constantPreDefine() {

	//iota是一个可被编译器修改的常量，每一个const出现时被置为0，没出现一次iota则值+1
	const ( //重置为0
		c0 = iota //c0=0
		c1 = iota //c1=1
		c2 = iota //c2=2
		c3        //与上面const赋值一样可以省略
	)

	const ( //重置为0
		a = 1 << iota //a=1
		b = 1 << iota //b=2
		c = 1 << iota //c=4
	)

	const ( //重置为0
		u         = iota * 42 //u=0
		v float64 = iota * 42 //v=42.0
		w         = iota * 42 //w=84
	)

	const x = iota //x=0,重置为0
	const y = iota //y=0
}
