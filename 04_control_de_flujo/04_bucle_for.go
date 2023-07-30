package main

import "fmt"

func main() {
		var i int
		for i <= 10 {
			fmt.Println(i)
			i++
		}

		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}

		// usa el break dentro de bucle for

		for i := 1; i <= 10; i++ {
			fmt.Println(i)
			if i == 5 {
				break
			}
		}

		// usa el continue dentro de bucle for

		for i := 1; i <= 10; i++ {
			if i == 5 {
				continue
			}
			fmt.Println(i)
		}
}
