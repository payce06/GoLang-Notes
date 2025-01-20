package main

import (
"fmt"
"sort"
)

func main() {
// Initialize a slice
list := []int{5, 2, 8, 1, 3}

// Print the original slice
fmt.Println("Original slice:", list)

// Access elements by index
fmt.Println("First element:", list[0]) // Access the first element
fmt.Println("Last element:", list[len(list)-1]) // Access the last element

// Append elements to the slice
list = append(list, 7, 9)
fmt.Println("After appending:", list)

// Create a slice from an existing slice (slicing)
subList := list[1:4] // Elements from index 1 to 3
fmt.Println("Sub-slice (1:4):", subList)

// Update elements
list[0] = 10
fmt.Println("After updating first element:", list)

// Iterate over the slice
fmt.Println("Iterating over the slice:")
for i, v := range list {
fmt.Printf("Index %d: Value %d\n", i, v)
}