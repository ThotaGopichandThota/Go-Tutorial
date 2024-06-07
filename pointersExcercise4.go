package main

import "fmt"

func main() {
	p := Person{
		firstName: "gopichand",
	}
	fmt.Println("value semantics:", p)
	p = changeName(p, "anil")
	fmt.Println(p)
	//pointer semantics
	changeNamePtr(&p, "chandu")
	fmt.Println("the pointer semantics:", p)
}

type Person struct {
	firstName string
}

func changeName(p Person, s string) Person {
	p.firstName = s
	return p
}
func changeNamePtr(p *Person, s string) {
	p.firstName = s
}
