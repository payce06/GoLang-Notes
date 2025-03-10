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