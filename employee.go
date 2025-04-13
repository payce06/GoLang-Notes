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

// Method to display employee information
func (e *Employee) DisplayEmployeeInfo() {
	fmt.Printf("Employee: %s\n", e.GetFullName())
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Salary: $%.2f\n", e.Salary)
}

// Method to increase the salary of an employee
func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary += amount
	fmt.Printf("Salary of %s increased by $%.2f. New Salary: $%.2f\n", e.GetFullName(), amount, e.Salary)
}

// Manager struct inherits from Employee
type Manager struct {
	Employee
	Department string
	Employees  []*Employee // Slice to hold the team of employees
}
// Method to add an employee to the manager's team
func (m *Manager) AddEmployee(employee *Employee) {
	m.Employees = append(m.Employees, employee)
	fmt.Printf("%s added to %s's team.\n", employee.GetFullName(), m.GetFullName())
}

// Method to display the manager's team
func (m *Manager) DisplayTeam() {
	fmt.Printf("%s's Team:\n", m.GetFullName())
	for _, employee := range m.Employees {
		employee.DisplayEmployeeInfo()
	}
}

// Override method to display manager's information
func (m *Manager) DisplayEmployeeInfo() {
	m.Employee.DisplayEmployeeInfo()
	fmt.Printf("Department: %s\n", m.Department)
}