package main

import (
	"fmt"
	"time"
)

type person struct {
	Name string
}

func main() {
	fmt.Println("== :", "Banana" == "Banana")
	fmt.Println("!= :", 3 != 2)
	fmt.Println("< :", 3 < 2)
	fmt.Println("> :", 3 > 2)
	fmt.Println("<= :", 3 <= 2)
	fmt.Println(">= :", 3 >= 2)

	dateOne := time.Unix(0, 0)
	dateTwo := time.Unix(0, 0)

	fmt.Println("data without pointer:", dateOne == dateTwo)
	fmt.Println("data with 'equals':", dateOne.Equal(dateTwo))

	personOne := person{Name: "João"}
	personTwo := person{Name: "João"}

	fmt.Println("Person:", personOne == personTwo)
}
