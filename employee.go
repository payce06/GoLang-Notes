package main

import (
	"fmt"
)

// Employee struct
type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float64
}

// Method to get full name of the employee
func (e *Employee) GetFullName() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}