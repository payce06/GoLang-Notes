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


// 5. Calculate the factorial of a number using a for loop
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
	result *= i
	}
	return result
	}
	
	// 6. Count occurrences of a specific character in a string using a for loop
	func countCharacter(s string, char rune) int {
	count := 0
	for _, c := range s {
	if c == char {
	count++
	}
	}
	return count
	}


// 7. Find the smallest number in an array using a for loop
func findSmallest(arr []int) int {
	smallest := arr[0]
	for _, num := range arr {
	if num < smallest {
	smallest = num
	}
	}
	return smallest
	}
	
	// 8. Multiply all the numbers in an array using a for loop
	func multiplyArray(arr []int) int {
	product := 1
	for _, num := range arr {
	product *= num
	}
	return product
	}


// 9. Print all prime numbers up to n using a for loop
func printPrimes(n int) []int {
	primes := []int{}
	for num := 2; num <= n; num++ {
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
	if num%i == 0 {
	isPrime = false
	break
	}
	}
	if isPrime {
	primes = append(primes, num)
	}
	}
	return primes
	}


// 10. Reverse the elements of an array using a for loop
func reverseArray(arr []int) []int {
	reversedArr := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
	reversedArr = append(reversedArr, arr[i])
	}
	return reversedArr
	}
	
	// 11. Concatenate all strings in an array using a for loop
	func concatenateStrings(arr []string) string {
	result := ""
	for _, str := range arr {
	result += str
	}
	return result
	}


// 12. Check if an array is sorted in ascending order using a for loop
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
	if arr[i] > arr[i+1] {
	return false
	}
	}
	return true
	}
	
	// 13. Find the length of the longest word in a string using a for loop
	func longestWordLength(s string) int {
	words := split(s, " ")
	maxLength := 0
	for _, word := range words {
	if len(word) > maxLength {
	maxLength = len(word)
	}
	}
	return maxLength
	}