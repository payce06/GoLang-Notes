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


// EnrollStudent method to enroll a student in a course
func (c *Course) EnrollStudent(student *Student) {
	for _, s := range c.Students {
		if s.StudentID == student.StudentID {
			fmt.Printf("%s is already enrolled in %s.\n", student.Name, c.CourseName)
			return
		}
	}
	c.Students = append(c.Students, student)
	fmt.Printf("%s has been enrolled in %s.\n", student.Name, c.CourseName)
}

// DisplayStudents method to display all students enrolled in the course
func (c *Course) DisplayStudents() {
	fmt.Printf("Course: %s (%s)\n", c.CourseName, c.CourseCode)
	fmt.Println("Enrolled Students:")
	for _, student := range c.Students {
		fmt.Printf("- %s (ID: %s)\n", student.Name, student.StudentID)
	}
}


// Student struct
type Student struct {
	Name      string
	Age       int
	StudentID string
	Courses   map[string][]float64 // Map to store courses and their grades
}