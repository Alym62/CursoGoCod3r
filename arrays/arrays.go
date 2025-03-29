package main

import "fmt"

func main() {
	// Arrays are of the same type and static
	var notes [3]float64
	notes[0], notes[1], notes[2] = 3, 9, 7
	fmt.Println(notes)

	total := 0.0
	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	media := total / float64(len(notes))
	fmt.Println("Média: ", media)
}
