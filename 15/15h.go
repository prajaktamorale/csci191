package main

import "fmt"

func main() {
	var number *[]int = new([]int)
	fmt.Println(number)
	fmt.Println(*number)
}
