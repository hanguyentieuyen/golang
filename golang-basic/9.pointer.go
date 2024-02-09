package main

import (
	"fmt"
)

type myStruct struct {
	number int
}

func main() {
	// var a int = 3
	// var b *int = &a // dùng & để lấy địa chỉ của con trỏ
	// fmt.Println(a, *b)
	// a = 45
	// fmt.Println(a, *b) // dùng * để lấy giá trị của con trỏ

	// sử dụng con trỏ với array, slice
	// x := []int{1, 2, 3}
	// y := x
	// fmt.Println(x, y)
	// x[1] = 8
	// fmt.Println(&x[1], &y[1])

	// sử dụng con trỏ với struct
	// var m *myStruct
	// fmt.Println(m)
	// //m = &myStruct{number: 89}
	// m = new(myStruct)
	// m.number = 20 // (*m).number = 20
	// fmt.Println(m)
	// fmt.Println(m.number) // fmt.Println((*m).number)

	// sử dụng con trỏ với map (giống slice)
	var q = map[string]int{"dog": 5, "cat": 2}
	p := q
	fmt.Println(q, p)
	p["dog"] = 3
	fmt.Println(q, p)
}
