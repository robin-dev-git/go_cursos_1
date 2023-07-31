package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("!Felicitaciones, adivinaste el número")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor.")
		} else {
			fmt.Println("El número es inválido")
		}
	}

	fmt.Println("Se acabaron los intentos.\nEl número era:", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Print("¿Quires jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)
	menu(eleccion)
}

func menu(eleccion string) string {
	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente.")
		jugarNuevamente()
	}
	return eleccion
}
