package main

import "fmt"

func main() {
    // ----- Basic if statement -----
    num := 10
    if num > 5 {
        fmt.Println("Basic if: num is greater than 5")
    }

    // ----- if-else statement -----
    if num%2 == 0 {
        fmt.Println("if-else: num is even")
    } else {
        fmt.Println("if-else: num is odd")
    }

    // ----- if-else if-else statement -----
    grade := 85
    if grade >= 90 {
        fmt.Println("if-else if-else: Grade is A")
    } else if grade >= 80 {
        fmt.Println("if-else if-else: Grade is B")
    } else if grade >= 70 {
        fmt.Println("if-else if-else: Grade is C")
    } else {
        fmt.Println("if-else if-else: Grade is F")
    }