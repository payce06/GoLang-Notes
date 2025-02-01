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

// Checking if a key exists
_, exists := map1["grapes"]
fmt.Println("Does 'grapes' exist in Map 1?", exists) // returns false if key doesn't exist

// Deleting a key-value pair from the map
delete(map1, "banana")
fmt.Println("Map 1 after deleting 'banana':", map1)

// Iterating over a map using a range loop
fmt.Println("\nIterating over Map 1:")
for key, value := range map1 {
fmt.Printf("Key: %s, Value: %d\n", key, value)
}

// Getting the length of a map (number of key-value pairs)
fmt.Println("\nLength of Map 1:", len(map1))

// Checking if a key exists using the comma ok idiom
age, ok := map2["age"]
if !ok {
fmt.Println("\nKey 'age' does not exist in Map 2.")
} else {
fmt.Println("Age:", age)
}

// Merging two maps (Manually copying entries from map2 to map1)
for key, value := range map2 {
map1[key] = len(value) // Example: Storing length of string values in map1
}
fmt.Println("\nMap 1 after merging with Map 2:", map1)

// Clearing a map (Setting it to an empty map)
map1 = make(map[string]int)
fmt.Println("\nMap 1 after clearing:", map1)

// Checking if the map is empty
if len(map1) == 0 {
fmt.Println("Map 1 is empty.")
} else {
fmt.Println("Map 1 is not empty.")
}
}
