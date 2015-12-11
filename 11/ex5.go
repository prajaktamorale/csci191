package main

import "fmt"

func main() {
	a := 2
	fmt.Println("A is:", a)
	fmt.Println("Memory Address of a is:", &a)
	var b *int = &a
	fmt.Println("b is a pointer to memory address of a:", b)
} 
