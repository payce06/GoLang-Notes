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


// 3. Function to find the second largest element in an array.
func secondLargest(array []int) int {
    largest := -1 << 31  // Set to minimum possible value
    secondLargest := -1 << 31
    for _, num := range array {
        if num > largest {
            secondLargest = largest
            largest = num
        } else if num > secondLargest && num != largest {
            secondLargest = num
        }
    }
    return secondLargest
}

// 4. Function to remove duplicates from an array.
func removeDuplicates(array []int) []int {
    uniqueArray := []int{}
    seen := map[int]bool{}
    for _, item := range array {
        if !seen[item] {
            seen[item] = true
            uniqueArray = append(uniqueArray, item)
        }
    }
    return uniqueArray
}

// 5. Function to reverse an array.
func reverseArray(array []int) []int {
    reversedArray := []int{}
    for i := len(array) - 1; i >= 0; i-- {
        reversedArray = append(reversedArray, array[i])
    }
    return reversedArray
}

// 6. Function to find the intersection of two arrays.
func arrayIntersection(array1, array2 []int) []int {
    intersection := []int{}
    seen := map[int]bool{}
    for _, item := range array1 {
        seen[item] = true
    }
    for _, item := range array2 {
        if seen[item] {
            intersection = append(intersection, item)
            seen[item] = false // Avoid duplicates
        }
    }
    return intersection
}