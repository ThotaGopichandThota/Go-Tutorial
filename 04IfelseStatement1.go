package main

import "fmt"

func main() {
	var num int

	fmt.Print("enter the num:")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("it is even number")

	} else {
		fmt.Println("it is odd number")
	}

}
