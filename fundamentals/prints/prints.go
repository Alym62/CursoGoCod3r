package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha. ")

	fmt.Println("Nova")
	fmt.Println("linha")

	x := 3.1415
	fmt.Println("[Println] - O valor de x é:", x)

	xs := fmt.Sprint(x)
	fmt.Println("[Println Sprint] - O valor de x é: " + xs)

	fmt.Printf("[Printf] - O valor de x é: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Ops!"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
