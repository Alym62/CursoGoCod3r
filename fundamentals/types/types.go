package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	/*
		Unsigned (positive) integers
		[uint8, uint16, uint32, uint64]
	*/
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	/*
		With sign (positive) integers
		[int8, int16, int32, int64]
	*/
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é:", i1)
	fmt.Println("O tipo de i1 é:", reflect.TypeOf(i1))

	// Represents a mapping of the Unicode table (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println(i2)

	/*
		Numbers reals
		[float32, float64]
	*/
	var x float32 = 49.99
	fmt.Println("O tipo x é:", reflect.TypeOf(x))

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é Aly"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da minha string é:", len(s1))

	// String with multiple rows
	s2 := `Olá
	meu
	nome
	é
	Aly
	`
	fmt.Println("O tamanho da string de multiplas linhas é:", len(s2))

	// Char? Not exists
	char := 'a'
	fmt.Println("O tipo char é:", reflect.TypeOf(char))
}
