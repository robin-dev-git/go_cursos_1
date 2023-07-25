package main

import (
	// "math"

	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	// fmt.Println(math.MinInt64, math.MaxInt64)
	// var valueBool boo = false

	fullName := "Robin Espejo \t(alias \"roelcode\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var r rune = '♥'
	fmt.Println(r)
}
