package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "Prajakta"
	fmt.Println("a: ", reflect.TypeOf(a))
	fmt.Printf( "a: %T\n", a)
}
