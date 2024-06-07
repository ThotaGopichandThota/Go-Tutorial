package main

import "fmt"

type Student struct {
	StdName       string
	StdRollNo     int
	StdBranch     string
	StdPercentage float64
}

func main() {
	fmt.Println("enter the student name:")
	var stdname string
	fmt.Scanln(&stdname)
	fmt.Println("enter the student role number:")
	var stdroleno int
	fmt.Scanln(&stdroleno)
	fmt.Println("enter the student branch")
	var stdbranch string
	fmt.Scanln(&stdbranch)
	fmt.Println("enter the student percentage:")
	var stdpercentage float64
	fmt.Scanln(&stdpercentage)
	std := Student{
		StdName:       stdname,
		StdRollNo:     stdroleno,
		StdBranch:     stdbranch,
		StdPercentage: stdpercentage,
	}
	pointer := &std
	fmt.Println(*pointer)
	//accessing the struct fields
	fmt.Println(pointer.StdName)
	fmt.Println(pointer.StdRollNo)
	fmt.Println(pointer.StdBranch)
	fmt.Println(pointer.StdPercentage)

	//modifying the student  details in struct fields by using the dereference pointer
	(*pointer).StdName = "chandu"
	(*pointer).StdRollNo = 123
	(*pointer).StdBranch = "civil"
	(*pointer).StdPercentage = 7.9
	fmt.Println("_____________after modifying the student details__________")
	fmt.Println((*pointer).StdName)
	fmt.Println((*pointer).StdRollNo)
	fmt.Println((*pointer).StdBranch)
	fmt.Println((*pointer).StdPercentage)
}
