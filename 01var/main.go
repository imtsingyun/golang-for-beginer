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

const pi = 3.14;

func main() {
	name = "tom"
	age = 19
	isOk = true
	
	fmt.Printf("name: %s, age: %d, isOk: %t\n", name, age, isOk)
	fmt.Printf("color: %s, weight: %f, height: %f\n", color, weight, height)
	
	var s1 string = "abc"
	var s2 = "efg";
	fmt.Printf("s1=%s, s2= %s\n", s1, s2)
	
	s3 := "xyx";
	fmt.Println(s3)
	fmt.Println(pi)	
}