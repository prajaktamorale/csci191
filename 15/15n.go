package main

import "fmt"

func main() {
	var number *[]int = new([]int)
	fmt.Println(*number)
	var point []int = make([]int, 2, 10)
	fmt.Println(point)
}

