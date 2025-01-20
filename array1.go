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

    fmt.Println("\nIterating with range:")
    for index, value := range arr3 {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

    // ----- 4. Using the len function -----
    fmt.Printf("\nLength of Array 1: %d\n", len(arr1))
    fmt.Printf("Length of Array 3: %d\n", len(arr3))

    // ----- 5. Copying arrays -----
    fmt.Println("\nCopying arrays:")
    copyArr := arr2 // Copies the entire array
    copyArr[0] = 999
    fmt.Println("Original Array 2:", arr2)
    fmt.Println("Copied Array (modified):", copyArr)

    // ----- 6. Multidimensional arrays -----
    fmt.Println("\nMultidimensional arrays:")
    var matrix [2][3]int // A 2x3 array
    matrix[0][0] = 1
    matrix[1][2] = 5
    fmt.Println("Matrix:", matrix)

    fmt.Println("Iterating over Matrix:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
        }
    }
}

