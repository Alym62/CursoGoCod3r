package main

import "fmt"

func buy(workOne, workTwo bool) (bool, bool, bool) {
	buyTv50 := workOne && workTwo // And
	buyTv32 := workOne != workTwo // Or (No exists OR exclusive in golang)

	buyIceCream := workOne || workTwo

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, ice := buy(true, true)
	fmt.Printf("Tv 50: %t, Tv 32: %t, Ice cream: %t, Health: %t", tv50, tv32, ice, !ice)
}
