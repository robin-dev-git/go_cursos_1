package main

import (
	"fmt"

	"rsc.io/quote"
)

// var firstName, lastName, age = "Robinson", "Espejo", 25

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	// Declaración de variables
	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName string = "Robin"
	// 	lastName  string = "Espe"
	// 	age       int    = 25
	// )

	// firstName = "Robinson"
	// lastName = "Espejo"
	// age = 25

	// var firstName, lastName, age = "Robinson", "Espejo", 25

	// firstName, lastName, age := "Robinson", "Espejo", 25

	firstName, lastName := "Robinson", "Espejo"
	age := 25

	fmt.Println(firstName, lastName, age)

}
