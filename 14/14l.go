package main

import ( 
		"fmt"
		"reflect"
)

func main() {
		d := "hello"
		g := 21
		
		fmt.Println(d, " is a ", reflect.TypeOf(d))
		fmt.Println(g, "is  ", reflect.TypeOf(g))
}
