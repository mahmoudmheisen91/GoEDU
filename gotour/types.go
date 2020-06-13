package main

import "fmt"

var print = fmt.Println

// factored decleration
var (
	test1 bool = true
	x     int  = 1 << 10
)

func main() {
	fmt.Printf("Type is %T, Value is %v\n", x, x)
	var a int
	var c string
	g := 0.867 + 0.5i // complex128
	fmt.Printf("%v %q\n", a, c)
}
