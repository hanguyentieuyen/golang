package main

import (
	"fmt"
	"strconv"
)

// Declare global variable
// global scope
var (
	n int    = 7
	m string = "ha"
)

// Export  global scope
var Number int = 55 // viết hoa

func main() {
	// block scope
	// 3 way declare variable
	var number int
	number = 44
	fmt.Println(number)

	var number1 int = 13
	fmt.Println(number1)

	number2 := 78
	fmt.Printf("%v, %T", number2, number2)

	// Type Inference
	var name = 4
	fmt.Println(name)
	fmt.Println(n)
	var n int = 9
	fmt.Println(n, m)

	// convert type
	// number -> number
	var i float64 = 3.5
	fmt.Printf("%v, %T", i, i)
	fmt.Println()
	//var j float32 = 3.14
	var j int = int(i)
	fmt.Printf("%v, %T", j, j)
	// number -> string
	var a int = 3
	fmt.Printf("%v, %T", a, a)
	fmt.Println()
	//var j float32 = 3.14
	var b string = strconv.Itoa(a) // nếu ko dùng package strconv thì convert string theo ascii
	fmt.Printf("%v, %T", b, b)
}
