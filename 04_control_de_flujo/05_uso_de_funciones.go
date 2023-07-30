package main

import "fmt"

func main() {
	saludo := hello("Robin")
	fmt.Println(saludo)

	fmt.Println(hello("Robin"))

	sum, mul := calc(4, 5)

	fmt.Println("La suma es: ", sum)
	fmt.Println("La multiplicación es: ", mul)


}

// usar funcion con imprimir en consola
/*
func hello(name string) {
	// fmt.Println("Hola desde la función")
	fmt.Println("hola,", name)
}
*/

// usar una funcion con retornar sin imprimir
func hello(name string) string {
	return "Hola, " + name
}

/*
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}
*/


func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return 
}