package main

import "fmt"

func main() {
	// i := 0
	// for {
	// 	if i%2 == 0 {
	// 		if i == 4 {
	// 			i++
	// 			continue // không thực hiện các câu lệnh bên dưới, chuyển sang vòng lặp tiếp theo
	// 		}
	// 		fmt.Println(i)
	// 	i++
	// 	if i > 10 {
	// 		break // thoát ra khỏi vòng lặp
	// 	}
	// }

	// Loop:
	// 	for a := 1; a <= 5; a++ {
	// 		// 1 vòng lặp a sẽ chạy hết tất cả vòng lặp b
	// 		for b := 1; b <= 5; b++ {
	// 			if a > 1 {
	// 				break Loop // thoát khỏi vòng lặp for đã được gán nhãn
	// 			}
	// 			fmt.Println(a * b)
	// 			if b > 3 {
	// 				break // thoát ra khỏi vòng lặp gần nhất
	// 			}
	// 		}
	// 	}

	// Duyệt array/slice
	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// Duyệt map
	m := map[string]int{
		"yen":   17,
		"lucia": 25,
		"jane":  40,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Duyệt chuỗi
	s := "Hello World"
	for index, value := range s {
		fmt.Println(index, string(value)) // convert ascii to string
	}
}
