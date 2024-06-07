package main

import "fmt"

func main() {
	fmt.Println("well come to pointers topic ")
	// normal initiization of variable
	x := 23
	fmt.Println("the value of x:", x)
	// taking of the pointer variable using var keyword without mentioning the type of variable
	var p = &x
	fmt.Println("the value stored in the variable:", x)
	fmt.Println("the address of the variable of x:", &x)
	fmt.Println("the value stored in the pointer varible of p:", *p)
	// assigning the new value to the pointer
	*p = 45
	fmt.Println("the value of x after assegning the new to the pointer:", x)
}
