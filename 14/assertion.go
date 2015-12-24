package main

import "fmt"


func main() {
	var exp interface{} = 7
	str, ok := exp.(string)
	if ok {
		fmt.Println("%T\n", str)
	} else {
		fmt.Println("not a string\n")
	}
}
