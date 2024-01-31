package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	age  int
	name string
}

type Student struct {
	Person
	number  int `Maximum can not over 100` // tag dùng để trả về thông tin cho user
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
	// studentInfor := Student{
	// 	number: 3,
	// 	name:   "Yen",
	// 	isMale: false,
	// 	age:    18,
	// 	subject: []string{
	// 		"Math",
	// 		"English",
	// 	},
	// }

	// Tao struct rỗng sau đó khai báo các kiểu dữ liệu
	// studentInfor := Student{}
	// studentInfor.name = "yen"
	// studentInfor.number = 3
	// studentInfor.age = 18
	// studentInfor.isMale = false
	// studentInfor.subject = []string{"Math", "Englist"}

	// copy struct
	// studentInfor := struct{ name string }{name: "yen"}
	// copyStudent := studentInfor
	// copyStudent.name = "hayen" // tương tự như array, copy struct tạo vùng nhớ mới
	// fmt.Println(copyStudent)

	//tag
	studentInfor := Student{}
	t := reflect.TypeOf(studentInfor)
	field, _ := t.FieldByName("number")
	fmt.Println(field)
	fmt.Println(studentInfor)
}
