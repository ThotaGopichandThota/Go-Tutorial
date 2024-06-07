package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "gopi",
		Age:  23,
	}
	p2 := Person{
		Name: "chandu",
		Age:  77,
	}
	p3 := Person{
		Name: "thota",
		Age:  24,
	}
	people := []Person{p1, p2, p3}
	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

}

type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}
func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
