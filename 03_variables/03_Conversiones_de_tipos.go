package main

import (
	// "math"
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	var integer16 int16
	var integer32 int32

	fmt.Println(int32(integer16) + integer32)
	
	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)
}
