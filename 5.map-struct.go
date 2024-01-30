package main

import "fmt"

type Student struct {
	number  int
	age     int
	name    string
	isMale  bool
	subject []string
}

func main() {
	// cấu trúc dữ liệu map: tâp hợp các cặp key value
	// studentNameAgeMap := make(map[string]int)
	// studentNameAgeMap = map[string]int{
	// 	"yen":  18,
	// 	"john": 25,
	// 	"mike": 35,
	// }
	// delete(studentNameAgeMap, "yen")
	// studentNameAgeMap["mike"] = 36
	// studentNameAgeMap["alice"] = 19

	// copyMap := studentNameAgeMap
	// copyMap["alice"] = 20
	// _, isExist := studentNameAgeMap["yen"]
	// fmt.Println(studentNameAgeMap)
	// fmt.Println(isExist)
	// fmt.Println(len(studentNameAgeMap))
	// fmt.Println(copyMap)

	// struct: tập hợp kiểu dữ liệu mới được tự định nghĩa
	studentInfor := Student{
		number: 3,
		name:   "Yen",
		isMale: false,
		age:    18,
		subject: []string{
			"Math",
			"English",
		},
	}

	fmt.Println(studentInfor)
}
