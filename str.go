package main

import (
"fmt"
"strings"
)

func main() {
// Initialize a string
str := "Hello, GoLang World!"

// Print the original string
fmt.Println("Original string:", str)

// Length of the string
fmt.Println("Length of the string:", len(str))

// Accessing individual characters (returns byte values)
fmt.Println("First character (byte):", str[0])  // Prints the byte value of 'H'

// Convert string to uppercase
upperStr := strings.ToUpper(str)
fmt.Println("Uppercase string:", upperStr)

// Convert string to lowercase
lowerStr := strings.ToLower(str)
fmt.Println("Lowercase string:", lowerStr)



// Check if string contains a substring
contains := strings.Contains(str, "GoLang")
fmt.Println("Contains 'GoLang'?", contains) // Returns true

// Check if string starts with a prefix
startsWith := strings.HasPrefix(str, "Hello")
fmt.Println("Starts with 'Hello'?", startsWith) // Returns true

// Check if string ends with a suffix
endsWith := strings.HasSuffix(str, "World!")
fmt.Println("Ends with 'World!'?", endsWith) // Returns true

// Finding the index of a substring
index := strings.Index(str, "GoLang")
fmt.Println("Index of 'GoLang':", index) // Returns index of first occurrence

// Replacing a substring
replacedStr := strings.ReplaceAll(str, "World", "Golang")
fmt.Println("String after replacement:", replacedStr)


// Splitting a string into a slice of substrings
splitStr := strings.Split(str, ", ")
fmt.Println("Split string:", splitStr) // Splits by ", "

// Trimming whitespace from both ends
trimmedStr := strings.TrimSpace("   Hello, GoLang!   ")
fmt.Println("Trimmed string:", trimmedStr)

// Removing leading whitespace
trimLeftStr := strings.TrimLeft("   Hello, GoLang!", " ")
fmt.Println("Trimmed left string:", trimLeftStr)

// Removing trailing whitespace
trimRightStr := strings.TrimRight("   Hello, GoLang!   ", " ")
fmt.Println("Trimmed right string:", trimRightStr)

// Repeat a string multiple times
repeatedStr := strings.Repeat("Go! ", 3)
fmt.Println("Repeated string:", repeatedStr) // Repeats "Go! " three times

// Checking if string is empty
isEmpty := str == ""
fmt.Println("Is the string empty?", isEmpty) // Returns false


// Joining elements of a string slice into a single string
words := []string{"Hello", "GoLang", "World"}
joinedStr := strings.Join(words, " ")
fmt.Println("Joined string:", joinedStr) // Joins with a space between words

// String formatting using fmt
name := "Paice"
age := 17
formattedStr := fmt.Sprintf("My name isS %s and I am %d years old.", name, age)
fmt.Println("Formatted string:", formattedStr)

// Checking for a specific character in a string (example: 'o')
charIndex := strings.IndexByte(str, 'o')
fmt.Println("Index of first occurrence of 'o':", charIndex)

// Counting occurrences of a character in a string
count := strings.Count(str, "o")
fmt.Println("Count of 'o' in string:", count)

// Comparing two strings
str1 := "GoLang"
str2 := "Golang"
comparison := strings.Compare(str1, str2)
fmt.Println("Comparison result:", comparison) // Returns 0 if equal, -1 if str1 < str2, 1 if str1 > str2

// Padding the string with a specific character to a certain length
paddedStr := fmt.Sprintf("%-20s", str)
fmt.Println("Padded string:", paddedStr)
}

package main

import (
"fmt"
"math"
)

// ----- 1. Defining a Struct -----
type Rectangle struct {
Length, Width float64
}

type Circle struct {
Radius float64
}

// ----- 2. Methods for Structs -----
func (r Rectangle) Area() float64 {
return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
return 2 * (r.Length + r.Width)
}

func (c Circle) Area() float64 {
return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
return 2 * math.Pi * c.Radius
}

// ----- 3. Interfaces -----
type Shape interface {
Area() float64
}


// ----- 4. Polymorphism with Interfaces -----
func printShapeDetails(s Shape) {
	fmt.Printf("Shape Area: %.2f\n", s.Area())
	}

	func main() {
	// ----- 5. Using Structs -----
	rect := Rectangle{Length: 5, Width: 3}
	circ := Circle{Radius: 4}

	fmt.Println("Rectangle Details:")
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	fmt.Println("\nCircle Details:")
	fmt.Printf("Area: %.2f\n", circ.Area())
	fmt.Printf("Circumference: %.2f\n", circ.Circumference())

	// ----- 6. Using Interfaces -----
	fmt.Println("\nUsing Interfaces:")
	printShapeDetails(rect)
	printShapeDetails(circ)

	// ----- 7. Slice of Interfaces -----
	fmt.Println("\nSlice of Interfaces:")
	shapes := []Shape{rect, circ}
	for _, shape := range shapes {
	printShapeDetails(shape)
	}
	}