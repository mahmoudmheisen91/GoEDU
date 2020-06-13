package main

import "fmt"

var print = fmt.Println

func add(x, y int) int {
	return x + y
}

func myfunc1(x, y int) (int, int) {
	a := x + y
	b := x - y
	return a, b
}

// named retyrn values
func myfunc2(x int) (y, z int) {
	y = 3
	z = 4
	return // naked return
}

func main() {
	print(add(4, 7))
	a, b := myfunc1(8, 2)
	print(myfunc1(4, 7))
	print(a, b)
	print(myfunc2(12))

	var i, j, k = 1, true, "str"
	print(i, j, k)
}
