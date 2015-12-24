package main

import "fmt"

func main() {
	f("Monday", "Tuesday", "Wednesday")
	f("Thursday", "Friday", "Saturday", "Sunday")
	bSlice := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	f(bSlice...)
	f()
}

func f(days ...string) {
	fmt.Println(days)
}

