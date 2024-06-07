package main

import "fmt"

func main() {
	temp := increment()
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
}
func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
