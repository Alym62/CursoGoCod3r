package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func numberRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	result(0.1)
	result(7.1)

	concept := noteToConcept(9.8)
	fmt.Println(concept)

	// IF with init
	if i := numberRandom(); i > 5 {
		fmt.Println("Ganhou!!!", i)
	} else {
		fmt.Println("Perdeu!!!", i)
	}
}
