package main

import "fmt"

func largest(numbers ...int) int {
	var greatest int
	for _, a := range numbers {
		if a > greatest {
			greatest = a
		}
	}
	return greatest
}

func main() {
	max := largest(1, 2, 3, 4, 5, 6, 7, 8, 9,10)
	fmt.Println(max)
}
