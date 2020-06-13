package main

import (
	"fmt"
	"math"
)

var print = fmt.Println

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	// for i := 0; i < 3; i++ {
	// 	print(i)
	// }

	// sum := 4
	// for sum < 7 {
	// 	print(sum)
	// 	sum++
	// }

	// print(sqrt(2), sqrt(-4))

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)
	// }

	defer fmt.Println("hi") // after the function is returned

	switch {
	case 1 > 1:
		print(1)
	case 1 < 1:
		print(2)
	case 1 == 1:
		print(3)
	}

}
