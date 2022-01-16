package ch2declarations

import "fmt"
// Boiling prints the boiling point of water.


// 1. const: declare a constant boilingF
const boilingF = 212.0

// 2. func: declare a function 
func Declaration() {
	// 3. var: declare a local variable f
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))
}