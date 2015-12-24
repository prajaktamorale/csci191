package main

import (
		"strconv"
		"fmt"
)

func main() {
	var p int = 60
	name := "Prajakta" + strconv.Itoa(p)
	f, _ := strconv.Atoi("50")
	fmt.Println(name)
	fmt.Println(f)
}
