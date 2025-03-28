package main

import "fmt"

func result(note float64) {
	if note >= 7 {
		fmt.Println("Aprovado com nota:", note)
	} else {
		fmt.Println("Reprovado com nota:", note)
	}
}

func main() {
	result(0.1)
	result(7.1)
}
