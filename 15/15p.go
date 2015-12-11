package main

import "fmt"

type customer struct {
	name string
	id int
}

func main() {
	c1 := customer{"Prajakta", 23}
	c2 := new(customer) //new can be used to add variables in structure
	c2.name = "Rama"
	c2.id = 45	
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1.name)
	c1.name = "Richa"
	fmt.Println(c1.name)
}
