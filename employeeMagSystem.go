package main

import (
	"fmt"
	"log"
)

func main() {
	Employee1 := addEmployee()
	log.Println("Employee", Employee1)
	employee2 := addEmployee()
	log.Println(employee2)
	Employee3 := addEmployee()
	log.println(Emlpoyee3)
	Employee2 := updateEmployee(2, "anil", 50000, "hr")
	log.Println("emloyee:", Employee2)
	Employees := displayAllRecords()
	log.Println("All employeees are: ", Employees)

}

type Employee struct {
	Id         int
	Name       string
	Salary     float32
	Department string
}

var employees []Employee

func addEmployee() Employee {
	fmt.Println("enter the employee id: ")
	var id int
	fmt.Scanln(&id)
	fmt.Println("enter the employee name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("enter the employee salary:")
	var salary float32
	fmt.Scanln(&salary)
	fmt.Println("enter the employee department: ")
	var department string
	fmt.Scanln(&department)

	emp := Employee{Id: id, Name: name, Salary: salary, Department: department}
	employees = append(employees, emp)
	fmt.Println("employee is added")
	return emp
}

// update the record  name and salary by employee id
func updateEmployee(id int, newName string, newSalary float32, newDepartment string) Employee {
	for i, emp := range employees {
		if emp.Id == id {
			employees[i].Name = newName
			employees[i].Salary = newSalary
			employees[i].Department = newDepartment
			fmt.Println(newName, newSalary, newDepartment)
			fmt.Println("record is updated")
			return employees[i]
		}

	}
	fmt.Println("record not found")
	return Employee{}
}
func displayAllRecords() []Employee {
	if len(employees) == 0 {
		fmt.Println("no records found")
		return nil
	}
	fmt.Println("employee records")
	for _, emp := range employees {
		fmt.Println(emp.Id, emp.Name, emp.Salary, emp.Department)
	}
	return employees
}
