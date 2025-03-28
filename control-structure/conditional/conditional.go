package main

import "fmt"

func noteToConcept(note float64) string {
	if note >= 9 && note <= 10 {
		return "A"
	} else if note >= 8 && note < 9 {
		return "B"
	} else if note >= 5 && note < 8 {
		return "C"
	} else {
		return "D"
	}
}

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

	concept := noteToConcept(9.8)
	fmt.Println(concept)
}
