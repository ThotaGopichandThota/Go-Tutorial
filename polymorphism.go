package main

import "fmt"

func main() {
	p1 := Person{
		name: "Gopichand",
	}
	sa := secreteAgent{
		Person: Person{
			name: "chandu",
		},

		ltk: true,
	}
	// p1.speak()
	// sa.speak()
	diplay(p1)
	diplay(sa)
}

type human interface {
	speak()
}

// polymorphism method for ineterface
func diplay(h human) {
	h.speak()
}

type Person struct {
	name string
}
type secreteAgent struct {
	Person
	ltk bool
}

func (p Person) speak() {
	fmt.Println("Iam ", p.name)
}
func (sa secreteAgent) secreteAgent() {
	fmt.Println("iam  a secreteAgent", sa.name)
}
