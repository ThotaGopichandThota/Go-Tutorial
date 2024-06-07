package main

import "fmt"

func main() {
	//value semantics
	num := 5
	fmt.Println(num)
	addOne(num)
	fmt.Println(num)

	//pointer semantics
	fmt.Println("pointer semantics")
	num1 := 5
	fmt.Println(num1)
	modify(&num1)
	fmt.Println(num1)
}

// value semantics
func addOne(v int) {
	v = 10
	fmt.Println(v)
}
func modify(x *int) {
	*x = 10
}
