package main

import (
	"fmt"
	"math"
)

// perimetro, hipotenusa,
func main() {
	var area, perimetro, hipotenusa, a, b float64

	fmt.Print("ingrese la longitud de lado \"A\" del triángulo rectángulo: ")
	fmt.Scanln(&a)
	fmt.Print("ingrese la longitud de lado \"B\" del triángulo rectángulo: ")
	fmt.Scanln(&b)

	hipotenusa = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	perimetro = a + b + hipotenusa

	area = (a * b) / 2

	fmt.Println("Resultados: \n1. Hipotensua: ", hipotenusa, "\n2. Perímetro: ", perimetro, "\n3. Área: ", area)

}
