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