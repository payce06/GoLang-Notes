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

// Department struct
type Department struct {
	Name     string
	Managers []*Manager // Slice to hold managers in the department
}

// Method to add a manager to the department
func (d *Department) AddManager(manager *Manager) {
	d.Managers = append(d.Managers, manager)
	fmt.Printf("%s added to %s department.\n", manager.GetFullName(), d.Name)
}

// Method to display department information
func (d *Department) DisplayDepartmentInfo() {
	fmt.Printf("Department: %s\n", d.Name)
	for _, manager := range d.Managers {
		manager.DisplayEmployeeInfo()
	}
}

// Main function
func main() {
	// Create employee objects
	employee1 := &Employee{"John", "Doe", 30, 50000}
	employee2 := &Employee{"Jane", "Smith", 28, 55000}
	employee3 := &Employee{"Alice", "Johnson", 35, 60000}

	// Create manager objects
	manager1 := &Manager{Employee: Employee{"Michael", "Brown", 40, 80000}, Department: "IT"}
	manager2 := &Manager{Employee: Employee{"Sara", "Dave", 38, 85000}, Department: "HR"}

	// Create department objects
	itDepartment := &Department{Name: "IT"}
	hrDepartment := &Department{Name: "HR"}

	// Add employees to managers
	manager1.AddEmployee(employee1)
	manager1.AddEmployee(employee2)
	manager2.AddEmployee(employee3)

	// Add managers to departments
	itDepartment.AddManager(manager1)
	hrDepartment.AddManager(manager2)

	// Display department details
	fmt.Println("\nIT Department:")
	itDepartment.DisplayDepartmentInfo()
	fmt.Println("\nHR Department:")
	hrDepartment.DisplayDepartmentInfo()

	// Increase salary of an employee
	employee1.IncreaseSalary(5000)

	// Display updated department details
	fmt.Println("\nUpdated IT Department:")
	itDepartment.DisplayDepartmentInfo()
}
