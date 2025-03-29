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

func noteToConceptChallenge(note float64) string {
	switch {
	case note >= 9 && note <= 10:
		return "A"
	case note >= 8 && note < 9:
		return "B"
	case note >= 5 && note < 8:
		return "C"
	default:
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

func noteToConceptWithSwitch(note float64) string {
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	default:
		return "D"
	}
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

	conceptWithSwitch := noteToConceptWithSwitch(10)
	fmt.Println(conceptWithSwitch)

	hour := time.Now()
	switch {
	case hour.Hour() < 12:
		fmt.Println("Bom dia")
	case hour.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	resultChallenge := noteToConceptChallenge(10)
	fmt.Println(resultChallenge)
}
