package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}

func (p *Persona) diHola(){
	fmt.Println("Hola, mi nombre es:", p.nombre)
}

func main() {
	/*
	var x int = 10
	var p *int = &x // puntero
	fmt.Println(&x)
	fmt.Println(p)
	*/
	/*
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)
	*/

	p := Persona{"robin", 24, "robin@gmail.com"}
	p.diHola()
}
/*
func editar(x *int)  {
	*x = 20
}
*/