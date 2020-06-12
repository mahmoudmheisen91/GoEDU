package main

import (
	"fmt"
	"math"
)

func main() {
	const theta = math.Pi / 6
	angle := math.Sin(theta)

	fmt.Println(angle)

	const d = 200 // has no type until its given one
	fmt.Println(int64(d))
}
