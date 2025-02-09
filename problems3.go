package main

import (
"fmt"
"math/rand"
"time"
)

// 1. Function to find the maximum number in an array
func findMaximum(arr []int) int {
maximum := arr[0]
for _, num := range arr {
if num > maximum {
maximum = num
}
}
return maximum
}

// 2. Function to remove duplicates from an array
func removeDuplicates(arr []int) []int {
unique := []int{}
seen := map[int]bool{}
for _, item := range arr {
if !seen[item] {
seen[item] = true
unique = append(unique, item)
}
}
return unique
}

// 3. Function that takes an array of numbers and returns a new array with each number squared
func squareElements(arr []int) []int {
squared := []int{}
for _, num := range arr {
squared = append(squared, num*num)
}
return squared
}

// 4. Function to sort an array of strings in alphabetical order
func sortStrings(arr []string) []string {
for i := 0; i < len(arr); i++ {
for j := i + 1; j < len(arr); j++ {
if arr[i] > arr[j] {
arr[i], arr[j] = arr[j], arr[i]
}
}
}
return arr
}


// 5. Function that returns the last element of an array
func lastElement(arr []int) int {
	if len(arr) == 0 {
	return -1 // Return -1 for empty arrays
	}
	return arr[len(arr)-1]
	}
	
	// 6. Function to find the median of an array of numbers
	func findMedian(arr []int) float64 {
	n := len(arr)
	if n == 0 {
	return 0
	}
	// Sort the array first
	for i := 0; i < n; i++ {
	for j := i + 1; j < n; j++ {
	if arr[i] > arr[j] {
	arr[i], arr[j] = arr[j], arr[i]
	}
	}
	}
	if n%2 == 0 {
	return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	return float64(arr[n/2])
	}