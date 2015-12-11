package main 

import "fmt"

const (
		_ = iota
		A = iota * 2
		B = iota * 2
		C = iota * 2
		D = iota * 2
		E = iota * 2
)

func main() {

	fmt.Println("Multiples of 2 upto 10:" , A, B, C, D, E)
}
