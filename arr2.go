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

// Check the length and capacity of the slice
fmt.Println("Length of the slice:", len(list))
fmt.Println("Capacity of the slice:", cap(list))

// Copy a slice into another slice
copySlice := make([]int, len(list)) // Create a new slice with the same length
copy(copySlice, list)
fmt.Println("Copied slice:", copySlice)

// Sort the slice (ascending order)
sort.Ints(list)
fmt.Println("Sorted slice:", list)

// Reverse a slice (custom logic)
for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
list[i], list[j] = list[j], list[i]
}
fmt.Println("Reversed slice:", list)

// Delete an element (e.g., element at index 2)
indexToRemove := 2
list = append(list[:indexToRemove], list[indexToRemove+1:]...)
fmt.Println("After deleting element at index 2:", list)

// Clear a slice (make it empty)
list = list[:0]
fmt.Println("After clearing the slice:", list)
}