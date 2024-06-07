package main

import "fmt"
// syntax for method
//  func (r receiver) identifier(p parameter(s)) (return(s)) { code }
func main() {
	p1 := Person{
		name: "Gopi",
		age:  23,
	}
	p2 := Person{
		name: "chandu",
		age:  23,
	}
	p1.speak()
	p2.speak()
}

type Person struct {
	name string
	age  int
}
//method
func (p Person) speak() {
	fmt.Println("Iam ", p.name, p.age)

}
