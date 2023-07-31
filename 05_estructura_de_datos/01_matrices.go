package main

import "fmt"

func main() {
	var a [5]int // array
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(a)

	var b = [5]int{10, 20, 30, 40, 50}
	a[0] = 100
	a[1] = 200
	fmt.Println(b)

	var c = [...]int{10, 20, 30, 40, 50}
	fmt.Println(c)
	fmt.Println(c[1])

	fmt.Println("---------------------")

	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	for index, value := range c {
		fmt.Printf("Índice: %d, valor: %d\n", index, value)
	}

	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matriz)
}
