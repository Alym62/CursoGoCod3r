package main

import (
	"fmt"
	"time"
)

func count(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	for count <= 30 {
		fmt.Println(count)
		count++
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// Infinite
	for {
		fmt.Println("Infinito")
		time.Sleep(time.Second * 5)
	}
}

func main() {
	count(10)
}
