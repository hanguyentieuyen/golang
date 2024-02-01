package main

import "fmt"

func main() {
	i := 0
	for {
		if i%2 == 0 {
			if i == 4 {
				i++
				continue // không thực hiện các câu lệnh bên dưới, chuyển sang vòng lặp tiếp theo
			}
			fmt.Println(i)
		i++
		if i > 10 {
			break // thoát ra khỏi vòng lặp
		}
	}
}
