package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)

    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    sum := num1 + num2

    fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}

