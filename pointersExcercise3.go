package main

import "fmt"

func main() {
	d1 := Dog{
		"anil",
	}
	youngRun(d1)
	d2 := Dog{
		"Gopi",
	}
	youngRun(d2)

	d2 = d2.changeName("chandu")
	youngRun(d2)
}

type Dog struct {
	Name string
}

func (d Dog) walk() {
	fmt.Println("my name is ", d.Name, "iam walking")
}
func (d Dog) run() {

	fmt.Println("my name is ", d.Name, "iam running")
}

type youngIn interface {
	run()
	walk()
}

func youngRun(y youngIn) {
	y.walk()
	y.run()
}
func (d Dog) changeName(s string) Dog {
	d.Name = s
	return d
}
