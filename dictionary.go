package main

import "fmt"

func main() {
    // ----- 1. Declaring and initializing maps -----
    // Declaring an empty map
    studentGrades := make(map[string]int)
    fmt.Println("Initial studentGrades map:", studentGrades)

    // Initializing a map with key-value pairs
    countryCapital := map[string]string{
        "USA":    "Washington, D.C.",
        "India":  "New Delhi",
        "Canada": "Ottawa",
    }
    fmt.Println("Country-Capital map:", countryCapital)

    // ----- 2. Accessing and modifying map elements -----
    studentGrades["Alice"] = 90
    studentGrades["Bob"] = 85
    fmt.Println("\nUpdated studentGrades map:", studentGrades)

    // Accessing a value by key
    grade := studentGrades["Alice"]
    fmt.Println("Alice's grade:", grade)

    // Modifying an existing value
    studentGrades["Alice"] = 95
    fmt.Println("Alice's updated grade:", studentGrades["Alice"])

    // ----- 3. Checking if a key exists -----
    score, exists := studentGrades["Charlie"]
    if exists {
        fmt.Println("Charlie's score:", score)
    } else {
        fmt.Println("Charlie is not in the studentGrades map.")
    }

    // ----- 4. Deleting elements from a map -----
    delete(studentGrades, "Bob")
    fmt.Println("\nAfter deleting Bob:", studentGrades)

    // ----- 5. Iterating over a map -----
    fmt.Println("\nIterating over countryCapital map:")
    for country, capital := range countryCapital {
        fmt.Printf("%s -> %s\n", country, capital)
    }

    // ----- 6. Map with different value types -----
    contactInfo := make(map[string]interface{})
    contactInfo["name"] = "John"
    contactInfo["age"] = 30
    contactInfo["isEmployed"] = true
    fmt.Println("\nContact info map with different value types:", contactInfo)
}