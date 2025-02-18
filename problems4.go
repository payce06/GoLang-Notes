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