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