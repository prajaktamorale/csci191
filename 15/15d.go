package main

import "fmt"

func main() {
	slice := []int{  1, 2, 3, 5, 4}
	myslice := []int{ 1, 2, 3, 4, 5, 6}
	slice = append(slice, myslice...)
	fmt.Println(slice)
}
