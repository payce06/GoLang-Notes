package main

import "fmt"

func main() {
    // Using the `var` keyword with explicit type
    var name string = "Payce"
    var age int = 16

    // Type inference
    var score = 100
    var isStudent = false

    // Short variable declaration
    greeting := "Hello, Go!"
    piValue := 3.14159

    // Declaring multiple variables
    var x, y, z int = 10, 20, 30
    a, b, c := "GoLang", 2.718, true

    // Constants
    const Pi = 3.14
    const Company = "The Boring Company"

    // Default zero values
    var uninitializedInt int
    var uninitializedString string
    var uninitializedBool bool
