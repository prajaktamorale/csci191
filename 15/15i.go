package main

import "fmt"

func main() {
	var name *string = new(string)
	fmt.Println(name)
	fmt.Println(*name)
}
