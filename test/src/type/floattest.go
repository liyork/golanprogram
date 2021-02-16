package main

import "math"

func testIntFloat() {

	var fvalue1 float32
	fvalue1 = 12
	fvalue2 := 12.0 //推导为float64=double
	//fvalue1 = fvalue2//float64不能赋值给float32
	fvalue1 = float32(fvalue2)
}

func isEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
