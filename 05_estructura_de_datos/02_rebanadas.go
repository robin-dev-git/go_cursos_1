package main

import "fmt"

func main() {
	/*
		diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

		diaRebanada := diasSemana[0:5]
		fmt.Println(diaRebanada)

		diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro dia")
		fmt.Println(diaRebanada)

		diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
		fmt.Println(diaRebanada)

		fmt.Println(len(diaRebanada))
		fmt.Println(cap(diaRebanada))
	*/
	/*
		nombre := make([]string, 5, 10)
		nombre[0] = "Robin"
		fmt.Println(nombre)
	*/
	rebanda1 := []int{1, 2, 3, 4, 5}
	rebanda2 := make([]int, 5)
/*
	copy(rebanda1, rebanda2)
	fmt.Println(rebanda1) //[0 0 0 0 0]
	fmt.Println(rebanda2) //[0 0 0 0 0]
	*/
	
	copy(rebanda2, rebanda1)

	fmt.Println(rebanda1) // [1 2 3 4 5]
	fmt.Println(rebanda2) // [1 2 3 4 5]
}
