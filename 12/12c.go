package main

import "fmt"

func main() {
    for i := 1; i <= 1000; i++ {
        if  i%15 == 0 {
        fmt.Println(i, "BuzzFizz")
        } else if i%3 == 0 {
		fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
		fmt.Println(i, "FizzBuzz")
		} else {
		fmt.Println(i)
    	}
	}
}
