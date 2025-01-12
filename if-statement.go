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

    // ----- Short variable declaration in if -----
    if age := 18; age >= 18 {
        fmt.Println("Short variable: You are eligible to vote")
    } else {
        fmt.Println("Short variable: You are not eligible to vote")
    }

    // ----- Nested if statements -----
    temperature := 30
    if temperature > 25 {
        if temperature > 35 {
            fmt.Println("Nested if: It's extremely hot")
        } else {
            fmt.Println("Nested if: It's warm")
        }
    } else {
        fmt.Println("Nested if: It's cool or cold")
    }

    // ----- Logical operators -----
    a := 5
    b := 15
    if a > 0 && b > 10 {
        fmt.Println("Logical operators: Both conditions are true")
    }

    if a > 0 || b < 10 {
        fmt.Println("Logical operators: At least one condition is true")
    }

    if !(a < 0) {
        fmt.Println("Logical operators: a is not negative")
    }
}