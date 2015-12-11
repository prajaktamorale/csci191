package main

import "fmt"

func main() {
	slice := make( []int, 5, 10)
	
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	fmt.Println(slice[1])
}
