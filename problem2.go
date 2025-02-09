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

// 7. Function to find the index of the first occurrence of an element in an array.
func findIndex(array []int, target int) int {
    for index, element := range array {
        if element == target {
            return index
        }
    }
    return -1 // Return -1 if the element is not found
}


// 8. Function to find all unique elements in an array.
func uniqueElements(array []int) []int {
    uniqueArray := []int{}
    seen := map[int]int{}
    for _, item := range array {
        seen[item]++
    }
    for item, count := range seen {
        if count == 1 {
            uniqueArray = append(uniqueArray, item)
        }
    }
    return uniqueArray
}


func main() {
    // Example usages
    fmt.Println(findLength([]int{1, 2, 3, 4})) // Output: 4
    fmt.Println(sumOfElements([]int{1, 2, 3, 4})) // Output: 10
    fmt.Println(secondLargest([]int{4, 1, 3, 2, 4})) // Output: 3
    fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 4})) // Output: [1 2 3 4]
    fmt.Println(reverseArray([]int{1, 2, 3, 4})) // Output: [4 3 2 1]
    fmt.Println(arrayIntersection([]int{1, 2, 3}, []int{2, 3, 4})) // Output: [2 3]
    fmt.Println(findIndex([]int{1, 2, 3, 4, 5}, 3)) // Output: 2
    fmt.Println(findIndex([]int{1, 2, 3, 4, 5}, 6)) // Output: -1
    fmt.Println(uniqueElements([]int{1, 2, 2, 3, 4, 4, 5})) // Output: [1 3 5]
}