package main

import "fmt"

func main() {
	d1 := Dogs{"anil"}
	d1.walk()
	d1.run()
	//youngIn(d1)

	d2 := &Dogs{"Gopi"}
	d2.walk()
	d2.run()
	youngIn(d2)
}

type Dogs struct {
	Name string
}

func (d Dogs) walk() {
	fmt.Println("my name is ", d.Name, "and iam walking")
}
func (d *Dogs) run() {
	d.Name = "rover"
	fmt.Println("my name is", d.Name, "and iam running")
}

type youngRun interface {
	walk()
	run()
}

func youngIn(y youngRun) {
	y.run()
}
