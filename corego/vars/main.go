package main

import "fmt"

func main() {
	//MAIN TYPES
	// string
	//bool
	//int
	//int int8 int16 int32 int64
	//unit uint8 uint16 unint32 unint64 unitptr
	//byte - alias for uint8
	//float32  float64
	//rune -alias for int32
	//complex64 complex128
	//using var keyword
	var name string = "Aman"
	//it is also valid
	var name1 = "Prem"
	var age = 37
	const isCool = true
	// this method can only be used inside  a function\
	style := "abc"
	size := 1.3
	fmt.Println(name, name1, age, isCool, style)
	fmt.Printf("%T\n%T\n%T\n", name, age, size)

}
