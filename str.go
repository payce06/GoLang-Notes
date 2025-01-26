package main

import (
"fmt"
"strings"
)

func main() {
// Initialize a string
str := "Hello, GoLang World!"

// Print the original string
fmt.Println("Original string:", str)

// Length of the string
fmt.Println("Length of the string:", len(str))

// Accessing individual characters (returns byte values)
fmt.Println("First character (byte):", str[0])  // Prints the byte value of 'H'

// Convert string to uppercase
upperStr := strings.ToUpper(str)
fmt.Println("Uppercase string:", upperStr)

// Convert string to lowercase
lowerStr := strings.ToLower(str)
fmt.Println("Lowercase string:", lowerStr)



// Check if string contains a substring
contains := strings.Contains(str, "GoLang")
fmt.Println("Contains 'GoLang'?", contains) // Returns true

// Check if string starts with a prefix
startsWith := strings.HasPrefix(str, "Hello")
fmt.Println("Starts with 'Hello'?", startsWith) // Returns true

// Check if string ends with a suffix
endsWith := strings.HasSuffix(str, "World!")
fmt.Println("Ends with 'World!'?", endsWith) // Returns true

// Finding the index of a substring
index := strings.Index(str, "GoLang")
fmt.Println("Index of 'GoLang':", index) // Returns index of first occurrence

// Replacing a substring
replacedStr := strings.ReplaceAll(str, "World", "Golang")
fmt.Println("String after replacement:", replacedStr)