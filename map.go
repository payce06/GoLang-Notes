package main

import "fmt"

func main() {
// Initialize a map using the make function
map1 := make(map[string]int)
map1["apple"] = 5
map1["banana"] = 3
map1["orange"] = 7

// Initialize a map with predefined values
map2 := map[string]string{
"firstName": "Alice",
"lastName":  "Johnson",
}

// Print the original map
fmt.Println("Map 1:", map1)
fmt.Println("Map 2:", map2)

// Accessing values by key
appleCount := map1["apple"]
fmt.Println("Apple count in Map 1:", appleCount)