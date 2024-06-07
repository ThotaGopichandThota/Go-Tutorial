/*
Create a type engine struct, and include this field
electric bool
Create a type vehicle struct, and include these fields
engine
make
model
doors
color
Create two VALUES of TYPE vehicle
use a composite literal
Print out each of these values.
Print out a single field from each of these values.
*/
package main

import "fmt"

func main() {
	v1 := vehicle{
		Engine: Engine{
			electric: false,
		},
		make:  "TaTa",
		model: "Innova",
		doors: 4,
		color: "Blue",
	}
	fmt.Println(v1)
	fmt.Println(v1.Engine.electric, v1.make, v1.model, v1.doors, v1.color)
	v2 := vehicle{
		Engine: Engine{
			electric: true,
		},
		make:  "TaTa",
		model: "nexon",
		doors: 4,
		color: "blue",
	}
	fmt.Println(v2)
	fmt.Println(v2.Engine.electric, v2.make, v2.model, v2.doors, v2.color)
}

type Engine struct {
	electric bool
}
type vehicle struct {
	Engine
	make  string
	model string
	doors int
	color string
}
