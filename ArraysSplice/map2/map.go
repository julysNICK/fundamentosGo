package main

import "fmt"

	func main() {
		funcSalarios := map[string]float64{
			"José João":      11325.45,
			"Gabriela Silva": 15456.78,
			"Pedro Junior":   1200.0,
		}

		funcSalarios["Rafael Filho"] = 1350.0
		delete(funcSalarios, "Inexistente")

		for nome, salario := range funcSalarios {
			fmt.Println(nome, salario)
		}

			
	}