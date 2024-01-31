package main

import (
	"fmt"
)

const (
	_ = 1 << iota // dịch bit <<
	red
	blue
	black
	white
)
const num = 13

func main() {
	const i int16 = 42
	const num = 31
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", blue, blue)
	fmt.Printf("%v, %T\n", black, black)

}

// 1. Định nghĩa hằng số
// 2. Khai báo hằng số
// 3. Đặc điểm của hằng số: hằng số bên trong block scope sẽ che (shadow)
// 						 lên hằng số global
// 4. kiểu enum: tập hợp nhiều hằng số
// 5. Tự động khởi tạo giá trị cho Enum bằng iota
