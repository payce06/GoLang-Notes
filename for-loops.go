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

    // ----- Using continue -----
    fmt.Println("\nUsing continue:")
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println("Odd number:", i)
    }

    // ----- Using range with slices -----
    fmt.Println("\nUsing range with a slice:")
    nums := []int{10, 20, 30, 40, 50}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // ----- Using range with maps -----
    fmt.Println("\nUsing range with a map:")
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    for key, value := range colors {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }

    // ----- Using range with a string -----
    fmt.Println("\nUsing range with a string:")
    for index, char := range "GoLang" {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

    // ----- Nested loops -----
    fmt.Println("\nNested loops:")
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}