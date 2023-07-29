package main

import (
	"fmt"
	"math"
	//"rsc.io/quote"
)

func main() {

	a := 10
	b := 3

	// suma
	fmt.Println(a + b)
	// resta
	fmt.Println(a - b)
	// multiplica
	fmt.Println(a * b)
	// Divido
	fmt.Println(a / b) // hay un error con el tipo de dato que no tiene numero decimal entonces...
	// solución
	// se converte en entero a flotante y ahora imprime
	fmt.Println(float64(a) / float64(b))
	// modulo
	fmt.Println(a % b)

	// decremento y incremento
	// incremento:
	fmt.Println("variable antes a y b ", a, b)
	a++
	b++
	fmt.Println("variable despues usa incremento a y b ", a, b)
	// decremento
	fmt.Println("variable antes a y b ", a, b)
	a--
	b--
	fmt.Println("variable despues usa decremento a y b ", a, b)

	// podemos utilizar asignación si mismo un variable se cambió algo
	a = 10
	a = a + 5
	fmt.Println(a)
	// sintaxis mas corto una linea, eso es lo mismo ante por ejemplo a = a + 5, puede ahorrar tiempo: a += 5.
	a += 5
	fmt.Println(a) // Puede utlizar asignación entre suma, resta, multiplicación, división y modulo

	// añadir un paquete con math
	// potencia
	c := math.Pow(2, 4)
	fmt.Printf("%.1f\n", c)

	// Raíz cuadrada
	d := math.Sqrt(4*4 + 2*2)
	fmt.Printf("%.1f\n", d)

	fmt.Println(math.Sqrt(64))

}
