package main

import "fmt"

func main() {
	fmt.Println("enter the user name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("enter the user mail id: ")
	var mail string
	fmt.Scanln(&mail)
	fmt.Println("enter the user age: ")
	var age int
	fmt.Scanln(&age)
	fmt.Println("enter the user Status")
	var status bool
	fmt.Scanln(&status)
	user := User{Name: name, Email: mail, Age: age, Status: status}

	fmt.Println("username:", user.Name)
	fmt.Println("user email Address: ", user.Email)
	fmt.Println("user age: ", user.Age)
	fmt.Println("user status ", user.Status)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
