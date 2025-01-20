package main

import "fmt"

func main() {
    // ----- 1. Declaring and initializing arrays -----
    var arr1 [5]int // Declares an array of 5 integers with default values (0)
    fmt.Println("Array 1 (default):", arr1)

    arr2 := [5]int{10, 20, 30, 40, 50} // Declares and initializes an array
    fmt.Println("Array 2 (initialized):", arr2)

    arr3 := [...]string{"Go", "is", "fun"} // Array size inferred from elements
    fmt.Println("Array 3 (inferred size):", arr3)

    // ----- 2. Accessing and modifying array elements -----
    arr1[0] = 100 // Assign a value to the first element
    arr1[4] = 500 // Assign a value to the last element
    fmt.Println("Array 1 (after modification):", arr1)

    fmt.Println("First element of Array 2:", arr2[0])
    fmt.Println("Last element of Array 2:", arr2[len(arr2)-1])

    // ----- 3. Iterating over arrays -----
    fmt.Println("\nIterating with for loop:")
    for i := 0; i < len(arr2); i++ {
        fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
    }