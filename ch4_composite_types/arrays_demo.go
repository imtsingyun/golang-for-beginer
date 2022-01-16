package ch4compositetypes

import "fmt"

func declareArrays() {
	var a [3]int				// array of 3 integers
	fmt.Println(a[0]) 			// print the first element
	fmt.Println(a[len(a) - 1]) 	// print the last element

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}