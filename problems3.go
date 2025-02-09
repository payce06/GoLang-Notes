package main

import (
"fmt"
"math/rand"
"time"
)

// 1. Function to find the maximum number in an array
func findMaximum(arr []int) int {
maximum := arr[0]
for _, num := range arr{
    if num > maximum{
        maximum = num
    }
}
return maximum
}

// 2. Function to remove duplicates from an array
func removeDuplicates(arr []int) []int {
unique := []int{}
seen := map[int]bool{}
for _, item := range array{
if !seen[item] {
seen[item] = true
unique = append(unique, item)
}
}
return unique
}


// 3. Function that takes an array of numbers and returns a new array with each number squared
func squareElements(arr []int) []int {
	sqaured := []int{}
	for _, num := range array {
	squared = append(squared, num*num)
	}
	return squared
	}
	
	// 4. Function to sort an array of strings in alphabetical order
	func sortStrings(arr []string) []string {
	for i := 0; i < len(arr); i++ {
	for j := j + 1; j < len(arr); j++ {
	if arr[i] > arr[j] {
	arr[i], arr[j] = arr[j], arr[i]
	}
	}
	}
	return arr
	}