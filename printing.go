package main

import "fmt"

func main() {
    // Basic printing
    fmt.Print("Hello, ")
    fmt.Println("World!") // Adds a newline at the end

    // Printing with formatting
    name := "Shreehar"
    age := 25
    height := 5.9
    fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

    // Using fmt.Sprintf to store formatted string
    formattedString := fmt.Sprintf("This is a formatted string: Name=%s, Age=%d", name, age)
    fmt.Println(formattedString)

    // Printing data types
    boolean := true
    character := 'A'
    fmt.Printf("Boolean: %t, Character: %c\n", boolean, character)

    // Printing arrays and slices
    numbers := []int{1, 2, 3}
    fmt.Printf("Numbers: %v\n", numbers)  // %v for default format
    fmt.Printf("Numbers with type: %#v\n", numbers) // %#v for Go syntax representation

    // Printing structs
    type Person struct {
        FirstName string
        LastName  string
    }
    person := Person{"Shreehar", "Joshi"}
    fmt.Printf("Person: %v\n", person)    // %v for default format
    fmt.Printf("Person with fields: %+v\n", person) // %+v to show field names
    fmt.Printf("Person Go syntax: %#v\n", person)   // %#v for Go syntax representation

    // Printing memory addresses
    fmt.Printf("Memory address of person: %p\n", &person)

    // Error printing
    fmt.Fprintf(fmt.Stderr, "This is an error message\n")
}
