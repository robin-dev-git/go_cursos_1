package main

import (
	"fmt"
<<<<<<< Updated upstream
	//"strconv"
=======

>>>>>>> Stashed changes
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

<<<<<<< Updated upstream
	// name := "Alex"
	// age := 28

	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	// fmt.Scanln(&name, &name2)
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)
	// fmt.Printf("Hola, me llamo %v y tengo %v años.\n", name, age)

	// greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	// fmt.Println(greeting)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de age es: %T\n", age)
=======
	fmt.Print("Otro mensaje")
>>>>>>> Stashed changes

}
