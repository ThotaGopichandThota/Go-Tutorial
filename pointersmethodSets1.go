package main

import "fmt"

func main() {
	d1 := Dogs{
		"anil",
	}
	d1.walk()
	d1.run()

	d2 := &Dogs{
		"Gopi",
	}
	d2.walk()
	d2.run()

}

type Dogs struct {
	Name string
}

func (d Dogs) walk() {
	fmt.Println("my name is ", d.Name, " and iam walking with my owner")
}

func (d *Dogs) run() {
	d.Name = "rover"
	fmt.Println("my name is ", d.Name, " and iam running with my owner")
}
