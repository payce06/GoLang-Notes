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

// EnrollInCourse method to enroll the student in a course
func (s *Student) EnrollInCourse(course *Course) {
	if _, exists := s.Courses[course.CourseCode]; !exists {
		s.Courses[course.CourseCode] = []float64{}
		course.EnrollStudent(s)
	} else {
		fmt.Printf("%s is already enrolled in %s.\n", s.Name, course.CourseName)
	}
}

// AddGrade method to add a grade for a specific course
func (s *Student) AddGrade(courseCode string, grade float64) {
	if grades, exists := s.Courses[courseCode]; exists {
		if grade >= 0 && grade <= 100 {
			s.Courses[courseCode] = append(grades, grade)
			fmt.Printf("Grade %.2f added for %s in course %s.\n", grade, s.Name, courseCode)
		} else {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
		}
	} else {
		fmt.Printf("%s is not enrolled in the course with code %s.\n", s.Name, courseCode)
	}
}