package main

import "fmt"

func main() {
	fmt.Println("enter the person name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("enter the person age")
	var age int
	fmt.Scanln(&age)
	fmt.Println("enter the person gender")
	var gender string
	fmt.Scanln(&gender)
	p := Person{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
	fmt.Println("initial person details:", p)
	pointer := &p
	modify(pointer)
	fmt.Println("modified person deatails:", p)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func modify(p *Person) {
	p.Name = "pawan kalyan"
	p.Age = 24
	p.Gender = "male"

}
