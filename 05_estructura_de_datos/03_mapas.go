package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	/*

		// fmt.Println(colors)
		fmt.Println(colors["rojo"])
		colors["negro"] = "#000000"

		fmt.Println(colors)

		// valor := colors["rojo"]
		// fmt.Println(valor)

		// valor, ok := colors["rojo"]
		// fmt.Println(valor, ok)

		if valor, ok := colors["rojo"]; ok {
			fmt.Println(valor)
		} else {
			fmt.Println("No existe la clave")
		}

		delete(colors, "azul")
		fmt.Println(colors)
	*/

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}

}
