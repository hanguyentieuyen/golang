/*
Primitive data types
1. Boolean
2. Numberíc
	Integer
			signed integer		int8  	int16 	int32 	int64
			unsigned integer	uint8 	uint16 	uint32
	Float
	Complex
3. Text
	String
	Byte
	Rune
*/

package main

import (
	"fmt"
)

func main() {
	// var a bool
	// a = 1 == 2
	// fmt.Printf("%v, %T\n", a, a)
	// fmt.Println()

	// var b uint8 = 255 // uint8: 0 -> 255
	// fmt.Printf("%v, %T\n", b, b)

	// var x int8 = 1
	// var y int8 = 9
	// fmt.Printf("%v, %T\n", x+y, x+y)
	// fmt.Printf("%v, %T\n", x-y, x-y)
	// fmt.Printf("%v, %T\n", x*y, x*y)
	// fmt.Printf("%v, %T\n", x/y, x/y)
	// fmt.Printf("%v, %T\n", x%y, x%y)

	// phép toán trên bit
	var m int8 = 10                    // 1010
	var n int8 = 3                     // 0011
	fmt.Printf("%v, %T\n", m&n, n&n)   // AND: 0010  => 2
	fmt.Printf("%v, %T\n", m|n, n|n)   // OR: 1011 => 11
	fmt.Printf("%v, %T\n", m^n, n^n)   // NOT: 1001 => 9
	fmt.Printf("%v, %T\n", m&^n, n&^n) // XOR: 0100
}
