package main

import "fmt"

func main() {

	fmt.Println("Enter a:")
	var a int
	fmt.Scanln(&a)
	fmt.Println("enter b:")
	var b int
	fmt.Scanln(&b)
	
	c := a + b
	fmt.Println("Result is", c)
} 

