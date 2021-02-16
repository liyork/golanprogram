package main

import "fmt"

//复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（ real），一个表示虚部（ imag）。
//math/cmplx
func testComplexNumber() {

	var value1 complex64 //由2个float32构成的复数类型
	value1 = 3.2 + 12i
	value2 := 3.2 + 12i        //value2是complex128类型
	value3 := complex(3.2, 12) //等同value2

	fmt.Printf("实部:%f,虚部:%f", real(value3), imag(value3))
}
