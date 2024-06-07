package main

import "fmt"

func main() {
	p1 := Person{
		FirstName: "Gopichand",
		LastName:  "thota",
		FavIce:    []string{"chocolate", "vennela"},
	}

	m := map[string]Person{
		p1.LastName: p1,
	}

	for _, value := range m {
		fmt.Println(value)
	}
}

type Person struct {
	FirstName string
	LastName  string
	FavIce    []string
}
