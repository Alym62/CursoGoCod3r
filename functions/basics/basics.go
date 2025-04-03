package main

import "fmt"

func showOneString(str string) {
	fmt.Println(str)
}

func showStrings(first, second string) {
	fmt.Println(first, second)
}

func returnString() string {
	return "Olá"
}

func multiplesReturns() (string, string) {
	return "Olá", "Mundo"
}

func main() {
	showOneString("Oi")
	showStrings("Aly", "Thai")
	fmt.Println(returnString())

	hello, world := multiplesReturns()
	fmt.Println(hello, world)
}
