
package main


import(
        "fmt"
)

func main() {

    fmt.Println("Enter A")
    var a int
    fmt.Scanln(&a)
    fmt.Println("Enter B")
    var b int
    fmt.Scanln(&b)
    c := a + b
    fmt.Println("Result:", c)
    }

