package main

import "fmt"

// create a value and assign it to a variable  and print the address of the variable
func main() {
	m := 23
	fmt.Println(m)
	fmt.Println(&m)
}
