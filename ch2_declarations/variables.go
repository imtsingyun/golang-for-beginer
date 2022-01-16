package ch2declarations

import "fmt"

// declare a variable: var name type = expression
func DeclareVariables() {
	var ss string;
	fmt.Println("ss = " + ss)
	
	var i, j, k int
	fmt.Printf("i = %d, j = %d, k = %d\n", i, j, k)
	var b, f, s = true, 2.3, "four"
	fmt.Printf("b = %t, f = %f, s = %s", b, f, s)
}