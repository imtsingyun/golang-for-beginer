package ch2declarations

import (
	"fmt"
	"math/rand"
)

// short variable declaration may be used to declare and initialize local variables within a function
// it takes the form `name := expression` 
// and the type of `name` is determined by the type of `expression`
func ShortVariable() {
	freq := rand.Float64() * 3.0
	fmt.Printf("freq = %f\n", freq)
}