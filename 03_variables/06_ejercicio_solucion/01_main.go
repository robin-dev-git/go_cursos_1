package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const presicion = 2

	// Entrada de dato
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado2)

	// Procesos
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	// Salida de datos
	fmt.Printf("\nÁrea: %.*f\n", presicion, area)
	fmt.Printf("Perímetro: %.*f\n", presicion, perimetro)
}
