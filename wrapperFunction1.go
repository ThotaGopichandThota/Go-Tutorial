package main

import "fmt"

func main() {
	fmt.Println("enter the first value:")
	var a int
	fmt.Scanln(&a)
	fmt.Println("enter the second value:")
	var b int
	fmt.Scanln(&b)
	sum := Addlog(a, b)
	fmt.Println("sum:", sum)
}
func add(a, b int) int {
	d := a + b
	return d
}

// wrapper function for logging

func Addlog(a, b int) int {
	fmt.Println("adding", a, "and", b)
	result := add(a, b)
	return result

}
