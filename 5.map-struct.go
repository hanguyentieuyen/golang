package main

import (
	"fmt"
)

func main() {
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int{
		"yen":  18,
		"john": 25,
		"mike": 35,
	}
	delete(studentNameAgeMap, "yen")
	studentNameAgeMap["mike"] = 36
	studentNameAgeMap["alice"] = 19
	_, isExist := studentNameAgeMap["yen"]
	fmt.Println(studentNameAgeMap)
	fmt.Println(isExist)
	fmt.Println(len(studentNameAgeMap))
}
