package main

import (
	"fmt"
	"strings"
)

// Course struct
type Course struct {
	CourseName string
	CourseCode string
	Students   []*Student // Slice to store enrolled students
}