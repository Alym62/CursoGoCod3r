package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Sub =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("Multi =>", a*b)
	fmt.Println("Mod =>", a%b)

	// Bitwise
	fmt.Println("AND =>", a&b)
	fmt.Println("OR =>", a|b)
	fmt.Println("XOR =>", a^b)

	// Use Math
	fmt.Println("Max =>", math.Max(2.0, 1.0))
	fmt.Println("Min =>", math.Min(2.0, 1.7))
}
