package main

import "fmt"

func main() {
	// Maps must be initialized
	approved := make(map[int]string)
	fmt.Println(approved)

	approved[1234567] = "João"
	approved[834782190] = "Pedro"
	approved[78237823] = "Ana"

	fmt.Println(approved)

	for tin, name := range approved {
		fmt.Printf("Nome: %s - CPF: %d\n", name, tin)
	}

	fmt.Println(approved[834782190])

	// Delete with key of map
	delete(approved, 834782190)
	fmt.Println(approved[834782190])

	// Working with maps
	salaries := map[string]float64{
		"Aly":     3850.95,
		"Thai":    2122.62,
		"Leiloca": 4234.98,
		"Papito":  15000,
	}

	salaries["Chefe"] = 30000
	delete(salaries, "Não existe nenhuma chave com esse nome!!!")

	for name, salary := range salaries {
		fmt.Printf("Nome: %s - R$ %.2f\n", name, salary)
	}

	// Map nested
	nestedMap := map[string]map[string]float64{
		"Dev Pleno": {
			"Aly": 6000,
		},
		"Vendedora": {
			"Thai": 3000,
		},
	}

	delete(nestedMap, "Dev PLeno")

	for funcs, names := range nestedMap {
		for name, salary := range names {
			fmt.Printf("Nome: %s (%s) - R$ %.2f\n", name, funcs, salary)
		}
	}
}
