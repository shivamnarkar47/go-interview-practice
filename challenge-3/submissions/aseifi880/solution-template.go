package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
    for i, emp := range m.Employees {
        if emp.ID  == id {
            m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
        }
    }
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
    totalSalary := 0.0
    numberOfEmployees := 0
	for _, e := range m.Employees {
	    totalSalary += e.Salary
	    numberOfEmployees += 1
	}
	if numberOfEmployees == 0 {
	    return 0.0
	}
	return totalSalary / float64(numberOfEmployees)
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	for _, e := range m.Employees {
	    if e.ID == id {
	        return &e
	    }
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
