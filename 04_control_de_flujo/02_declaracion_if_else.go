package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	if condicion {
		...
	} else {
		...
	}
	*/
/*
	t := time.Now()

	hora := t.Hour()

	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
*/
	// sintaxis diferentes 
	//pero ese el programa es lo mismo 
	//lo que imprime uno de sintaxis son iguales
	
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
}
