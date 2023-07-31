package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
/*
	var p Persona
	p.nombre = "Robin"
	p.edad = 25
	p.correo = "robin@gmail.com"

	fmt.Println(p)
*/

	p := Persona {"Robin", 25, "robin@gmail.com"}
	fmt.Println(p)
	fmt.Println(p.nombre)

	p.edad = 26
	fmt.Println(p)

	p2 := Persona {"Juan", 30, "juan@gmail.com"}
	fmt.Println(p2)



}
