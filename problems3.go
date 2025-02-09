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

// 9. Function that shuffles the elements of an array
func shuffleArray(arr []int) []int {
	// Use the current time as a seed for random number generation
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]int, len(arr))
	copy(shuffled, arr)
	rand.Shuffle(len(shuffled), func(i, j int) {
	shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
	}
	
	// 10. Function to rotate the elements of an array to the right by k steps
	func rotateArray(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
	return arr
	}
	k %= n
	return append(arr[n-k:], arr[:n-k]...)
	}

// 11. Function to split an array into chunks of a specific size
func chunkArray(arr []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(arr); i += size {
	end := i + size
	if end > len(arr) {
	end = len(arr)
	}
	chunks = append(chunks, arr[i:end])
	}
	return chunks
	}
	
	// 12. Function that takes an array and a number and returns a new array with elements greater than the given number
	func filterGreater(arr []int, num int) []int {
	var result []int
	for _, item := range arr {
	if item > num {
	result = append(result, item)
	}
	}
	return result
	}


// 13. Function to remove falsy values from an array
func removeFalsy(arr []interface{}) []interface{} {
	var result []interface{}
	for _, item := range arr {
	if item != nil && item != false && item != 0 && item != "" {
	result = append(result, item)
	}
	}
	return result
	}
	
	// 14. Function that finds the intersection of two arrays
	func intersection(arr1, arr2 []int) []int {
	seen := map[int]bool{}
	var result []int
	for _, item := range arr1 {
	seen[item] = true
	}
	for _, item := range arr2 {
	if seen[item] {
	result = append(result, item)
	seen[item] = false // Avoid duplicates
	}
	}
	return result
	}

// 15. Function to find the difference between two arrays
func difference(arr1, arr2 []int) []int {
	var result []int
	seen := map[int]bool{}
	for _, item := range arr2 {
	seen[item] = true
	}
	for _, item := range arr1 {
	if !seen[item] {
	result = append(result, item)
	}
	}
	seen = map[int]bool{}
	for _, item := range arr1 {
	seen[item] = true
	}
	for _, item := range arr2 {
	if !seen[item] {
	result = append(result, item)
	}
	}
	return result
	}

// 16. Function to concatenate two arrays
func concatenateArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
	}
	
	// 17. Function that returns the first n elements of an array
	func firstNElements(arr []int, n int) []int {
	if n > len(arr) {
	n = len(arr)
	}
	return arr[:n]
	}
	
	// 18. Function to find the common elements in three arrays
	func commonInThree(arr1, arr2, arr3 []int) []int {
	seen := map[int]int{}
	var result []int
	for _, item := range arr1 {
	seen[item]++
	}
	for _, item := range arr2 {
	seen[item]++
	}
	for _, item := range arr3 {
	seen[item]++
	}
	for item, count := range seen {
	if count == 3 {
	result = append(result, item)
	}
	}
	return result
	}

	
// 19. Function that returns the unique elements in an array
func uniqueElements(arr []int) []int {
	var unique []int
	seen := map[int]bool{}
	for _, item := range arr {
	if !seen[item] {
	seen[item] = true
	unique = append(unique, item)
	}
	}
	return unique
	}
	
	// 20. Function that merges two sorted arrays into one sorted array
	func mergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	var result []int
	for i < len(arr1) && j < len(arr2) {
	if arr1[i] < arr2[j] {
	result = append(result, arr1[i])
	i++
	} else {
	result = append(result, arr2[j])
	j++
	}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
	}