package main

import "fmt"

func main() {
	fmt.Println("enter the number:")
	var n int
	fmt.Scanln(&n)
	fmt.Println("the factorial of the given number using loop is:", factorialLoop(n))
	fmt.Println("the factorial of the given number is :", factorial(n))
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func factorialLoop(n int) int {
	total := 1
	for n > 0 {
		total = n * total
		n--
	}
	return total
}
