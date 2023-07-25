package main

import (
	// "math"

	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)
}
