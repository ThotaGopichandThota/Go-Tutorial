package main

import "fmt"

func main() {
	fmt.Println("enter the name of the person:")
	var name string
	fmt.Scanln(&name)
	fmt.Println(" enter the age of the person:")
	var age int
	fmt.Scanln(&age)
	p := Person{
		Name: name,
		Age:  age,
	}
	// pointer to the struct
	pointer := &p

	fmt.Println(*pointer)
	// updating the value of the struct

	//pointer.Name = "anil"
	fmt.Println("enter the modified person name:")
	fmt.Scanln(&pointer.Name)
	//pointer.Age = 24
	fmt.Println("enter the modified person age:")
	fmt.Scanln(&pointer.Age)
	fmt.Println(*pointer)
}

type Person struct {
	Name string
	Age  int
}
