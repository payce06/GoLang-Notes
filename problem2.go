package main

import "fmt"

// 1. Function to find the length of an array without using len().
func findLength(array []int) int {
    count := 0
    for range array {
        count++
    }
    return count
}

// 2. Function to find the sum of all elements in an array.
func sumOfElements(array []int) int {
    total := 0
    for _, num := range array {
        total += num
    }
    return total
}