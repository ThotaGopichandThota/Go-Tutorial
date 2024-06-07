package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age int
	var favIce []string

	// Get name input
	fmt.Println("Enter name:")
	fmt.Scanln(&name)

	// Get age input
	fmt.Println("Enter age:")
	fmt.Scanln(&age)

	// Get favorite ice cream flavors input
	fmt.Println("Enter favorite ice cream flavors (comma-separated):")
	var icecream string
	fmt.Scanln(&icecream)
	// Split the comma-separated string into slice of strings
	for _, flavor := range flavors(icecream) {
		favIce = append(favIce, flavor)
	}
	p1 := Person{
		Name:   name,
		Age:    age,
		FavIce: favIce,
	}

	for _, v := range p1.FavIce {
		fmt.Println(p1.Name, v)
	}
	fmt.Println(p1)
	fmt.Println(p1.FavIce)
}

func flavors(s string) []string {
	var result []string
	for _, str := range strings.Split(s, ",") {
		result = append(result, strings.TrimSpace(str))
	}
	return result
}

type Person struct {
	Name   string
	Age    int
	FavIce []string
}

