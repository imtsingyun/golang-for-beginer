package main

import "fmt"

// 声明变量
var name string // ""
var age int		// 0
var isOk bool	// false

// 批量声明
var (
	color string
	weight float32
	height float32
)

func main() {
	name = "tom"
	age = 19
	isOk = true
	
	fmt.Printf("name: %s, age: %d, isOk: %t", name, age, isOk)
	fmt.Printf("color: %s, weight: %f, height: %f", color, weight, height)
	
}