package main

import "fmt"

func main() {

	/// if expression condition
	m := map[string]int{
		"yen": 20,
	}

	if age, isExist := m["yen"]; isExist {
		fmt.Println(age)
	}

	age := 13

	if age > 6 && age < 11 {
		fmt.Println("Đủ tuổi học cấp 1")
	} else if age < 15 && age >= 11 {
		fmt.Println("Đủ tuổi học cấp 2")
	} else {
		fmt.Println(age)
	}

	// switch expression condition

	number := 2

	switch {
	case number <= 3:
		fmt.Println("Nhỏ hơn hoặc bằng 3")
		fallthrough // tiếp tục kiểm tra case bên dưới, khi case này đã thỏa điêu kiện
	case number <= 5:
		fmt.Println("Nhỏ hơn hoặc bằng 5")
		fmt.Println("Nhỏ hơn hoặc bằng 5")
		//break // ko thực hiện dòng bên dưới, thoát khỏi switch case
		fmt.Println("Nhỏ hơn hoặc bằng 5")
		fallthrough
	default:
		fmt.Println("Không thuộc trường hợp trên")
	}
}
