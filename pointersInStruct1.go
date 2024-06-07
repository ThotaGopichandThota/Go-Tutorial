package main

import "fmt"

func main() {
	fmt.Println("enter the employee name:")
	var empName string
	fmt.Scanln(&empName)
	fmt.Println("enter the employee id :")
	var empId int
	fmt.Scanln(&empId)
	fmt.Println("enter the employee salary :")
	var empSalary float64
	fmt.Scanln(&empSalary)
	fmt.Println("enter the employee Role:")
	var empRole string
	fmt.Scanln(&empRole)
	emp := Employee{
		EmpName:   empName,
		EmpId:     empId,
		EmpSalary: empSalary,
		EmpRole:   empRole,
	}
	pointer := &emp
	fmt.Println(*pointer)

	//accessing the fields of the struct one by one
	fmt.Println(pointer.EmpName)
	fmt.Println(pointer.EmpId)
	fmt.Println(pointer.EmpSalary)
	fmt.Println(pointer.EmpRole)

	//updating the values of the fields like name and salary of the struct
	fmt.Println("struct before updating the values:", *pointer)
	pointer.EmpName = "Gopichand"
	pointer.EmpSalary = 50000

	fmt.Println("struct after updating the values by using pointer", *pointer)
}

type Employee struct {
	EmpName   string
	EmpId     int
	EmpSalary float64
	EmpRole   string
}
