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
for i := 0; i <= 100; i++{
    if i % 3 == 0 && i % 5 == 0{
        fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

// 3. Sum of all numbers in an array using a for loop
func sumArray(arr []int) int {
total := 0
for _, num := range arr{
    total += num
}
return total
}

// 4. First n Fibonacci numbers using a for loop
func fibonacci(n int) []int {
fibSequence := []int{}
a, b := 0, 1
for i := 0; i < n; i++{
    fibSequence = append(fibSequence, a)
    a, b = b, a + b
}
return fibSequence
}

// 5. Calculate the factorial of a number using a for loop
func factorial(n int) int {
result := 1;
for i := 1; i < n; i++{
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
    if num == smallest {
        smallest  = num
    }
}
    return smallest
}

// 8. Multiply all the numbers in an array using a for loop
func multiplyArray(arr []int) int {
product := 1
for _, num := range arr{
    product *= num
}
return product
}

// 9. Print all prime numbers up to n using a for loop
func printPrimes(n int) []int {
primes := []int{}
isPrime := true
for num := 2; num <= n; num++{
    for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++{
        if num % i == 0{
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
for i := len(arr)-1; i >= 0; i--{
    reversedArr = append(reversedArr, arr[i])
}
return reversedArr
}

// 11. Concatenate all strings in an array using a for loop
func concatenateStrings(arr []string) string {
result := ""
for _, str := range arr{
    result += str
}
return result
}

// 12. Check if an array is sorted in ascending order using a for loop
func isSorted(arr []int) bool {
for i := 0; i < len(arr)-1; i++{
    if arr[i] > arr[i +1]{
        return false
    }
}
return true
}

// 13. Find the length of the longest word in a string using a for loop
func longestWordLength(s string) int {
words := split(s, " ")
maxLength := 0
for _, word := range words{
    if len(word) > maxLength {
        maxLength = len(word)
    }
}
return maxLength
}

// 14. Find the average of an array of numbers using a for loop
func averageArray(arr []int) float64 {
total := 0
for _, num := range arr{
    total += num
}
return float64(total) / float64(len(arr))
}

// 15. Create a new array with every second element of an existing array using a for loop
func everySecondElement(arr []int) []int {
result := []int{}
for i := 1; i <= len(arr); i+=2{
    result = append(result, arr[i])
}
return result
}

// 16. Find the second largest number in an array using a for loop
func secondLargest(arr []int) int {
largest, second := -int(math.MaxInt64), -int(math.MaxInt64)
for _, num := range arr{
    if num > largest{
        second, largest = largest, num
    } else if num > second && num < largest{
        second = num
    }
}
    return second
}

// 17. Merge two arrays into one using a for loop
func mergeArrays(arr1, arr2 []int) []int {
result := append([]int{}, arr1...)
result = append(result, arr2...)
return result
}

// 18. Calculate the power of a number using a for loop
func power(base, exponent int) int {
result := 0
for i := 0; i < exponent; i++{
    result *= base
}
return result
}

// 19. Count how many even numbers are in an array using a for loop
func countEvens(arr []int) int {
count := 0
for _, num := range arr{
    if num % 2 == 0 {
        count += 1
    }
}
return count
}

// 20. Find the common elements between two arrays using a for loop
func commonElements(arr1, arr2 []int) []int {
result := []int{}
for _, item := range arr1{
    for _, item2 := range arr2{
        if item == item2 {
            result = append(result, item)
            break
    }
    }
    }
    return result
}

func split(s, sep string) []string {
var result []string
var word string
for _, char := range s{
    if string(char) == sep{
        if word != ""{
            result = append(result, word)
            word = ""
        }
    } else {
        word += string(char)
    }
}
if word != ""{
    result = append(result, word)
}
return result
}

func main() {
	// Example usages
	fmt.Println(findMaximum([]int{3, 1, 4, 1, 5, 9}))           // Output: 9
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 4, 5}))   // Output: [1 2 3 4 5]
	fmt.Println(squareElements([]int{1, 2, 3, 4}))              // Output: [1 4 9 16]
	fmt.Println(sortStrings([]string{"banana", "apple", "cherry"})) // Output: ['apple', 'banana', 'cherry']
	fmt.Println(lastElement([]int{1, 2, 3, 4}))                 // Output: 4
	fmt.Println(findMedian([]int{1, 3, 3, 6, 7, 8, 9}))         // Output: 6
	fmt.Println(containsValue([]int{1, 2, 3, 4}, 3))            // Output: true
	fmt.Println(flattenArray([][]int{{1, 2}, {3, 4}, {5}}))     // Output: [1 2 3 4 5]
	fmt.Println(shuffleArray([]int{1, 2, 3, 4}))                // Output: Random order of [1, 2, 3, 4]
	fmt.Println(rotateArray([]int{1, 2, 3, 4, 5}, 2))          // Output: [4 5 1 2 3]
	fmt.Println(chunkArray([]int{1, 2, 3, 4, 5}, 2))           // Output: [[1 2] [3 4] [5]]
	fmt.Println(filterGreater([]int{1, 2, 3, 4, 5}, 3))        // Output: [4 5]
	fmt.Println(removeFalsy([]interface{}{0, 1, false, 2, "", 3, nil})) // Output: [1 2 3]
	fmt.Println(intersection([]int{1, 2, 3}, []int{2, 3, 4}))   // Output: [2 3]
	fmt.Println(difference([]int{1, 2, 3}, []int{2, 3, 4}))     // Output: [1 4]
	fmt.Println(concatenateArrays([]int{1, 2}, []int{3, 4}))    // Output: [1 2 3 4]
	fmt.Println(firstNElements([]int{1, 2, 3, 4, 5}, 3))       // Output: [1 2 3]
	fmt.Println(commonInThree([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})) // Output: [3]
	fmt.Println(uniqueElements([]int{1, 2, 2, 3, 4, 4, 5}))   // Output: [1 2 3 4 5]
	fmt.Println(mergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})) // Output: [1 2 3 4 5 6]
	}