package main

import "fmt"

func main() {
    // ----- Basic for loop -----
    fmt.Println("Basic for loop:")
    for i := 0; i < 5; i++ {
        fmt.Println("i =", i)
    }

    // ----- for loop as a while loop -----
    fmt.Println("\nFor as while loop:")
    count := 0
    for count < 3 {
        fmt.Println("count =", count)
        count++
    }

    // ----- Infinite loop with break -----
    fmt.Println("\nInfinite loop with break:")
    sum := 0
    for {
        sum++
        if sum >= 5 {
            fmt.Println("Breaking infinite loop, sum =", sum)
            break
        }
    }