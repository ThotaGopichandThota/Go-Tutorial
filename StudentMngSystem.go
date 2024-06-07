package main

import (
	"fmt"
	"log"
)

func main() {
	student1 := addStudent()
	log.Println(student1)
	student2 := addStudent()
	log.Println(student2)
	student := updateStd(2, "chandu", 79.9, "ece", "Gec")
	log.Println(student)
	display := displayStds()
	log.Println(display)

}

var students []Student

type Student struct {
	Id         int
	Name       string
	Percentage float32
	Stream     string
	ClgName    string
}

func addStudent() Student {
	fmt.Println("Enter the student id:")
	var id int
	fmt.Scanln(&id)
	fmt.Println("Enter the student name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Enter the student percentage:")
	var per float32
	fmt.Scanln(&per)
	fmt.Println("Enter the student stream:")
	var stream string
	fmt.Scanln(&stream)
	fmt.Println("Enter the student college name:")
	var clgName string
	fmt.Scanln(&clgName)
	std := Student{Id: id, Name: name, Percentage: per, Stream: stream, ClgName: clgName}
	students = append(students, std)
	fmt.Println("Record added")
	return std
}

func updateStd(id int, newName string, newPer float32, newStream string, newClg string) Student {
	for i, std := range students {
		if std.Id == id {
			students[i].Name = newName
			students[i].Percentage = newPer
			students[i].Stream = newStream
			students[i].ClgName = newClg

			fmt.Println("Record updated")
			return students[i]
		}
	}
	fmt.Println("No records found")
	return Student{}
}

func displayStds() []Student {
	if len(students) == 0 {
		fmt.Println("No record found")
		return nil
	}
	fmt.Println("Student records:")
	for _, std := range students {
		fmt.Println(std.Id, std.Name, std.Percentage, std.Stream, std.ClgName)
	}
	return students
}
