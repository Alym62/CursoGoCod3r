package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2 // type (float64)

	// type inference
	area := PI * m.Pow(radius, 2)
	fmt.Println("A área da circuferência é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Ops!"

	fmt.Println(g, h, i)
}
