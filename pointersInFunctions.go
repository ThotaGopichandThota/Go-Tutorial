package main

import "fmt"

func main() {
	//normal initilizing the variable

	var x int = 45
	fmt.Println(" the value of x before calling pointer function:", x)
	// taking the pointer variable and assigning the address of the x to it
	var pa *int = &x
	fmt.Println("the address of the x:", pa)
	//calling the pointer function
	pointer(pa)
	fmt.Println(" the value of the x after the pointer function calling:", x)
}
func pointer(a *int) {
	//dereferencing
	*a = 100

}
