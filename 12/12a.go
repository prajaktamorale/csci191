package main

import(
        "fmt"
)

func main() {

    fmt.Println("Enter A")
    var inputa int
    fmt.Scanln(&inputa)
    fmt.Println(inputa)
    fmt.Println("Enter B")
    var inputb int
    fmt.Scanln(&inputb)
    fmt.Println(inputb)
    c := inputa % inputb
    fmt.Println("Result:", c)
    }

