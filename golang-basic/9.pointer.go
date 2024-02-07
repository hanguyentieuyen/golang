package main

import (
	"fmt"
)

func main() {
	// var a int = 3
	// var b *int = &a // dùng & để lấy địa chỉ của con trỏ
	// fmt.Println(a, *b)
	// a = 45
	// fmt.Println(a, *b) // dùng * để lấy giá trị của con trỏ

	//
	x := []int{1, 2, 3}
	y := x
	fmt.Println(x, y)
	x[1] = 8
	fmt.Println(&x[1], &y[1])
}
