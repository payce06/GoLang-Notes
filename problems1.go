package main

import "fmt"

// 1. Function to check if a number is positive, negative, or zero.
func checkNumber(num int) string {
    if num > 0 {
        return "Positive"
    } else if num < 0 {
        return "Negative"
    }
    return "Zero"
}

// 2. Function to find the largest of three numbers.
func findLargestOfThree(a, b, c int) int {
    if a > b && a > c {
        return a
    } else if b > c {
        return b
    }
    return c
}

// 3. Function to determine if a year is a leap year.
func isLeapYear(year int) bool {
    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        return true
    }
    return false
}

// 4. Function to find the sum of all numbers in a slice (list in Go).
func sumOfList(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 5. Function to print all even numbers in a range.
func evenNumbersInRange(start, end int) []int {
    evens := []int{}
    for num := start; num <= end; num++ {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    return evens
}

// 6. Function to count the occurrences of a specific element in a slice.
func countOccurrences(numbers []int, target int) int {
    count := 0
    for _, num := range numbers {
        if num == target {
            count++
        }
    }
    return count
}

// 7. Function to find the smallest number in a slice.
func findSmallest(numbers []int) int {
    smallest := numbers[0]
    for _, num := range numbers {
        if num < smallest {
            smallest = num
        }
    }
    return smallest
}

// 8. Function to print the multiplication table of a number.
func multiplicationTable(n int) []string {
    table := []string{}
    for i := 1; i <= 10; i++ {
        table = append(table, fmt.Sprintf("%d x %d = %d", n, i, n*i))
    }
    return table
}

// 9. Function to calculate the sum of all odd numbers in a range.
func sumOfOddNumbers(start, end int) int {
    total := 0
    for num := start; num <= end; num++ {
        if num%2 != 0 {
            total += num
        }
    }
    return total
}

// 10. Function to check if a number is prime.
func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    // Example usages
    fmt.Println(checkNumber(-5))                  // Output: Negative
    fmt.Println(findLargestOfThree(3, 7, 5))      // Output: 7
    fmt.Println(isLeapYear(2024))                 // Output: true
    fmt.Println(sumOfList([]int{1, 2, 3, 4, 5}))  // Output: 15
    fmt.Println(evenNumbersInRange(1, 10))        // Output: [2 4 6 8 10]
    fmt.Println(countOccurrences([]int{1, 2, 2, 3, 4, 2}, 2))  // Output: 3
    fmt.Println(findSmallest([]int{3, 1, 4, 1, 5})) // Output: 1

    // Print multiplication table of 5
    for _, line := range multiplicationTable(5) {
        fmt.Println(line)
    }

    fmt.Println(sumOfOddNumbers(1, 10))          // Output: 25
    fmt.Println(isPrime(17))                     // Output: true
}