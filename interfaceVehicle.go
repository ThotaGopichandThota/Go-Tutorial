package main

import "fmt"

func main() {
	fmt.Println("Enter the bike name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter the bike speed:")
	var speed int
	fmt.Scanln(&speed)

	bike := Bike{Name: name, BikeSpeed: speed}
	Details(bike)
}

type Vehicle interface {
	Speed() int
	details()
}

type Bike struct {
	Name      string
	BikeSpeed int
}

func (b Bike) Speed() int {
	return b.BikeSpeed
}

func (b Bike) details() {
	fmt.Printf("Bike name: %s, Speed: %d\n", b.Name, b.BikeSpeed)
}

func Details(v Vehicle) {
	v.details()
}
