package main

import "fmt"

func main() {
	temp := gopi()
	fmt.Println(temp)

	temp1 := chandu()
	fmt.Println(temp1())
}

// function return string argument
func gopi() string {
	return ("my name is Gopichand")
}

// function return function
func chandu() func() string {
	return func() string {
		return "iam from Andhra Pradesh"
	}
}
