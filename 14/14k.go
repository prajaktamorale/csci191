package main

import (
		"fmt"
		"math"
)

func main() {
	fmt.Println("Enter the value of x")
	var x float64
	fmt.Scanln(&x)
	y := int(math.Ceil(x))
	fmt.Println("Using Ceil the value of x is-", y)
}

