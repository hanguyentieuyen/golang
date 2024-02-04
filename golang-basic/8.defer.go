package main

import "fmt"

// defer dùng để điều khiển luồng, cho phép hoãn lại đến khi chạy hết các lệnh còn lại
// khi gặp defer thì câu lệnh được đưa vào stack - FILO
// Khi gọi lệnh defer thì biến trong câu lệnh sẽ là giá trị tại điểm đó

// panic dùng để xử lỗi ko mong muốn (exception)

func main() {
	// fmt.Println("Line 1")
	// defer fmt.Println("Line 2")
	// fmt.Println("Line 3")
	// a := 1
	// defer fmt.Println(a)
	// a = 33

	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("Create panic")
	// Khi gặp defer đưa vào stack
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	// Khi gặp panic kiểm tra trong stack và thực hiện defer
	panic("Lỗi không mong muốn xảy ra")
	fmt.Println("hello word") // Dòng dưới panic không đc thực thi.
}
