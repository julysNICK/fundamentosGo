package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma =>", a+b)
	fmt.Println("subtração", a-b)
	fmt.Println("divisão", a/b)
	fmt.Println("multiplicação" , a*b)
	fmt.Println("módulo", a%b)

	// bitwise
	
	fmt.Println("e", a&b)
	fmt.Println("ou", a|b)
	fmt.Println("xor", a^b)

	c := 3.0
	d := 2.0

	// outras operacoes usando math...
	fmt.Println("maior", math.Max(float64(a), float64(b)))
	fmt.Println("menor", math.Min(c, d))
	fmt.Println("exponenciação", math.Pow(c, d))
	
}