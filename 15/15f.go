package main

import "fmt"

func main() {
	slice := make([]int, 2, 2)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)
}
