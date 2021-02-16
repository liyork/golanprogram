package main

import "fmt"

func main() {
	fval := 110.48
	ival := 200
	sval := "This is a string. "
	fmt.Println("The value of fval is", fval)
	fmt.Printf("fval=%f, ival=%d, sval=%s\n", fval, ival, sval)
	fmt.Printf("fval=%v, ival=%v, sval=%v\n", fval, ival, sval)
}
