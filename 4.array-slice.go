package main

import "fmt"

func main() {
	// array
	// var newArr [3]int
	// newArr[1] = 10
	// points := [...]int{1, 2, 3, 4}
	// fmt.Printf("%v,%T\n", newArr[1], newArr)
	// fmt.Printf("%v,%T\n", points, points)
	// fmt.Println(len(points))

	// a := [3]int{2, 5, 10}
	// b := &a
	// b[1] = 20
	// fmt.Println(a)
	// fmt.Println(b)

	// slice: bản chất là array, đóng vai trò là con trỏ đầu cuối, truy xuất cùng địa chỉ
	// len: số lượng phần tử
	// capacity: khả năng duyệt tối đa, cap >= len, capacity phụ thuộc vào con trỏ cuối
	// toán tử ":" dùng để thay đổi con trỏ đầu, con trỏ cuối của 1 slice
	// phân biệt len và capacity của slice
	// m := []int{1, 2, 3, 4, 33, 45}
	// n := m[:]
	// o := m[3:]
	// p := m[:4]
	// q := m[3:6]
	// q[1] = 26
	// fmt.Printf("m %v, %v, %v\n", m, len(m), cap(m))
	// fmt.Printf("n %v, %v, %v\n", n, len(n), cap(n))
	// fmt.Printf("o %v, %v, %v\n", o, len(o), cap(o))
	// fmt.Printf("p %v, %v, %v\n", p, len(p), cap(p))
	// fmt.Printf("q %v, %v, %v\n", q, len(q), cap(q))

	// make(), append()
	// x := make([]int, 0)
	// //fmt.Printf("make:%v, %v, %v\n", x, len(x), cap(x))
	// x = append(x, 1)
	// //fmt.Printf("make:%v, %v, %v\n", x, len(x), cap(x))
	// x = append(x, 2, 3, 4)
	// fmt.Printf("make:%v, %v, %v\n", x, len(x), cap(x))
	// x = append(x, []int{5, 6, 7}...)
	// fmt.Printf("make:%v, %v, %v\n", x, len(x), cap(x))

	// FIFO stack
	e := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v, %v, %v\n", e, len(e), cap(e))

	f := e[1:]
	fmt.Printf("%v, %v, %v\n", f, len(f), cap(f))

	h := append(e[:2], e[3:]...)
	fmt.Printf("%v, %v, %v\n", h, len(h), cap(h))
}
