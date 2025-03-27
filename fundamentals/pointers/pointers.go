package main

import "fmt"

func main() {
	i := 10

	// Go no have arithmetic of pointer
	var pointer *int = nil
	pointer = &i // address pointer

	*pointer++ // getting pointer value

	fmt.Println(*pointer)
}
