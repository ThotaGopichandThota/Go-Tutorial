package main

import "fmt"

func main() {
	p1 := Person{
		Name:   "Gopichand",
		Age:    23,
		FavIce: []string{"chocolate", "vennela", "butterScotch"},
	}

	fmt.Println(p1)
	for _, v := range p1.FavIce {
		fmt.Println(p1.Name, "favorite Ice-cream is", v)
	}
	fmt.Println(p1.FavIce)

}

type Person struct {
	Name   string
	Age    int
	FavIce []string
}
