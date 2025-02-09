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