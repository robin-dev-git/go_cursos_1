package main

import (
	"fmt"

	"rsc.io/quote"
)

const Pi = 3.14

const (
	x = 100
	y = 0b100 // binario
	z = 0o12  // Ocatal
	w = 0xFF  // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	firstName, lastName := "Robinson", "Espejo"
	age := 25

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)

}
