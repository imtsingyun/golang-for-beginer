package main

import (
	"example/ch1_helloworld"
	ch2declarations "example/ch2_declarations"
	"fmt"
)

func main() {
	// Example1: ch1_HelloWorld
	fmt.Println("---------------------Example1: ch1_helloworld---------------------------")
	helloworld.SayHello()
	// Example2: ch2_declaration
	fmt.Println("---------------------Example2: ch2_declaration--------------------------")
	ch2declarations.Declaration()
	// Example3: ch2_declaretion_variables 
	fmt.Println("---------------------Example3: ch2_declaration--------------------------")
	ch2declarations.DeclareVariables()
}