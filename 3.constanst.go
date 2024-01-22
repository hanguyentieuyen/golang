package main

import (
	"fmt"
)

const (
	red   = 0
	blue  = 1
	black = 2
	white = 0
)
const num = 13

func main() {
	const i int16 = 42
	const num = 31
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", red, red)
}

// 1. Định nghĩa hằng số
// 2. Khai báo hằng số
// 3. Đặc điểm của hằng số: hằng số bên trong block scope sẽ che (shadow)
// 						 lên hằng số global
// 4. kiểu enum: tập hợp nhiều hằng số
