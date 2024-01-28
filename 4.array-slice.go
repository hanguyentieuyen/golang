package main

import (
	"fmt"
)

func main() {
	// array
	var newArr [3]int
	newArr[1] = 10
	points := [...]int{1, 2, 3, 4}
	fmt.Printf("%v,%T\n", newArr[1], newArr)
	fmt.Printf("%v,%T\n", points, points)
	fmt.Println(len(points))

	a := [3]int{2, 5, 10}
	b := &a
	b[1] = 20
	fmt.Println(a)
	fmt.Println(b)

	// slice
}
