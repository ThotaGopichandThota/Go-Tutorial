package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter the person name is: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("enter the person age is: ")
	var age int
	fmt.Scanln(&age)
	fmt.Println("enter the person gender is:")
	var gender string
	fmt.Scanln(&gender)
	person1 := Person{Name: name, Age: age, Gender: gender}
	fmt.Println(person1)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) String() string {
	// name := p.Name
	// age := strconv.Itoa(int(p.Age))
	// gender := p.Gender

	a:= fmt.Sprintf("Name: %s, Age: %d, Gender: %s", p.Name, p.Age, p.Gender)
	fmt.Println(a)
	return a
}
