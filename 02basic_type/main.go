package main

import "fmt"

func main() {
	n := 109
	fmt.Printf("%b\n", n) // binary 2
	fmt.Printf("%d\n", n) // decimal 10
	fmt.Printf("%o\n", n) // Octal 8
	fmt.Printf("%x\n", n) // hex 16  lower case 
	fmt.Printf("%X\n", n) // hex 16  capital

	fmt.Printf("%p\n", &n) // memory address 0xc0000aa058
}