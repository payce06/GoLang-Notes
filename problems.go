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

// 4. Function to check if a string is a palindrome.
func isPalindrome(s string) bool {
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
if s[i] != s[j] {
return false
}
}
return true
}

// 5. Function to calculate the nth Fibonacci number.
func fibonacci(n int) int {
if n <= 0 {
fmt.Println("Invalid input")
return -1
}
if n == 1 {
return 0
}
if n == 2 {
return 1
}
return fibonacci(n-1) + fibonacci(n-2)
}

// 6. Function to find the largest number in a list.
func findLargest(numbers []int) int {
largest := numbers[0]
for _, num := range numbers {
if num > largest {
largest = num
}
}
return largest
}

// 7. Function to reverse a string.
func reverseString(s string) string {
runes := []rune(s)
for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
runes[i], runes[j] = runes[j], runes[i]
}
return string(runes)
}

// 8. Function to count the number of vowels in a string.
func countVowels(s string) int {
vowels := "aeiouAEIOU"
count := 0
for _, ch := range s {
if strings.ContainsRune(vowels, ch) {
count++
}
}
return count
}

// 9. Function to convert Celsius to Fahrenheit.
func celsiusToFahrenheit(celsius float64) float64 {
return (celsius * 9 / 5) + 32
}

// 10. Function to find the GCD of two numbers.
func gcd(a, b int) int {
for b != 0 {
a, b = b, a%b
}
return a
}

func main() {
// Examples
fmt.Println(addNumbers(3, 5))                // Output: 8
fmt.Println(isEvenOrOdd(7))                  // Output: Odd
fmt.Println(factorial(5))                    // Output: 120
fmt.Println(isPalindrome("radar"))           // Output: true
fmt.Println(fibonacci(7))                    // Output: 8
fmt.Println(findLargest([]int{1, 3, 7, 0, 5})) // Output: 7
fmt.Println(reverseString("hello"))          // Output: "olleh"
fmt.Println(countVowels("Python Programming")) // Output: 4
fmt.Println(celsiusToFahrenheit(25))         // Output: 77.0
fmt.Println(gcd(48, 18))                     // Output: 6
}