/*
Create and use an anonymous struct with these fields:
first     string
friends   map[string]int
favDrinks []string
Print things.
*/
package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Gopi",
		friends: map[string]int{
			"anil":    25,
			"raj":     27,
			"subbu":   24,
			"sumanth": 24,
		},
		favDrinks: []string{
			"water",
			"alcohol",
		},
	}
	for i, v := range p1.friends {
		fmt.Println(p1.first, "-friends-", i, v)
	}
	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "-drinks-", v)
	}
}
