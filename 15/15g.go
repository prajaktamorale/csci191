package main

import "fmt"

func main() {
	mymap := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	mymap[4] = "fourth"
	mymap[5] = "fifth"
	fmt.Println("map with added entries: ", mymap)
	mymap[2] = "2nd"
	fmt.Println("map with updated entry: ", mymap)
	delete(mymap, 3)
	fmt.Println("map with deleted entry: ", mymap)
	fmt.Println(len(mymap))
	
	if val, exists := mymap[2]; exists {
		delete(mymap, 2)
		fmt.Println("value", val)
		fmt.Println("exists", exists)
	} else {
		fmt.Println("Value does not exist")
		fmt.Println("value", val)
		fmt.Println("exists", exists)
	}
	for key, val := range mymap {
		fmt.Println(key,  "--", val)
	}
}

