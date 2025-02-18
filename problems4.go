package main

import (
"fmt"
"math"
)

// 1. Reverse a string using a for loop
func reverseString(s string) string {
reversedStr := ""
for _, char := range s {
reversedStr = string(char) + reversedStr
}
return reversedStr
}

// 2. FizzBuzz for numbers 1 to 100
func fizzBuzz() {
for i := 1; i <= 100; i++ {
if i%3 == 0 && i%5 == 0 {
fmt.Println("FizzBuzz")
} else if i%3 == 0 {
fmt.Println("Fizz")
} else if i%5 == 0 {
fmt.Println("Buzz")
} else {
fmt.Println(i)
}
}
}

// 3. Sum of all numbers in an array using a for loop
func sumArray(arr []int) int {
	total := 0
	for _, num := range arr {
	total += num
	}
	return total
	}
	
	// 4. First n Fibonacci numbers using a for loop
	func fibonacci(n int) []int {
	fibSequence := []int{}
	a, b := 0, 1
	for i := 0; i < n; i++ {
	fibSequence = append(fibSequence, a)
	a, b = b, a+b
	}
	return fibSequence
	}