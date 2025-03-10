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

// CalculateGPA method to calculate and return the GPA
func (s *Student) CalculateGPA() float64 {
	totalGrades := 0.0
	totalCourses := 0
	for _, grades := range s.Courses {
		for _, grade := range grades {
			totalGrades += grade
			totalCourses++
		}
	}
	if totalCourses == 0 {
		return 0
	}
	return totalGrades / float64(totalCourses)
}

// DisplayInfo method to display the student's information
func (s *Student) DisplayInfo() {
	fmt.Printf("Student ID: %s\n", s.StudentID)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Println("Courses and Grades:")
	for courseCode, grades := range s.Courses {
		avg := 0.0
		for _, grade := range grades {
			avg += grade
		}
		avg = avg / float64(len(grades))
		fmt.Printf("  %s: %v (Average: %.2f)\n", courseCode, grades, avg)
	}
	fmt.Printf("GPA: %.2f\n", s.CalculateGPA())
}

// StudentManagementSystem struct
type StudentManagementSystem struct {
	Students map[string]*Student
	Courses  map[string]*Course
}

// AddStudent method to add a student to the system
func (sms *StudentManagementSystem) AddStudent(student *Student) {
	if _, exists := sms.Students[student.StudentID]; !exists {
		sms.Students[student.StudentID] = student
		fmt.Printf("Student %s added to the system.\n", student.Name)
	} else {
		fmt.Printf("Student ID %s already exists.\n", student.StudentID)
	}
}
// AddCourse method to add a course to the system
func (sms *StudentManagementSystem) AddCourse(course *Course) {
	if _, exists := sms.Courses[course.CourseCode]; !exists {
		sms.Courses[course.CourseCode] = course
		fmt.Printf("Course %s added to the system.\n", course.CourseName)
	} else {
		fmt.Printf("Course Code %s already exists.\n", course.CourseCode)
	}
}
// SearchStudent method to search for a student by their ID
func (sms *StudentManagementSystem) SearchStudent(studentID string) *Student {
	return sms.Students[studentID]
}

// DisplayAllStudents method to display all students in the system
func (sms *StudentManagementSystem) DisplayAllStudents() {
	for _, student := range sms.Students {
		student.DisplayInfo()
	}
}

// Main function
func main() {
	// Create the management system
	sms := &StudentManagementSystem{
		Students: make(map[string]*Student),
		Courses:  make(map[string]*Course),
	}