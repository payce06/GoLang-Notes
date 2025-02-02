package main

import (
"fmt"
"strings"
)

// 1. Function to add two numbers.
func addNumbers(a, b int) int {
return a + b
}

// 2. Function to check if a number is even or odd.
func isEvenOrOdd(number int) string {
if number%2 == 0 {
return "Even"
}
return "Odd"
}

// 3. Function to find the factorial of a number.
func factorial(n int) int {
if n == 0 || n == 1 {
return 1
}
return n * factorial(n-1)
}