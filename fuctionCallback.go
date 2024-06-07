package main

import "fmt"

func main() {
	temp := doMath(33, 44, add)
	fmt.Println(temp)
	temp1 := doMath(44, 33, subtract)
	fmt.Println(temp1)
}
func add(a int, b int) int {

	c := a + b
	return c
}
func subtract(a int, b int) int {
	d := a - b
	return d
}
func doMath(a int, b int, f func(int, int) int) int {
	k := f(a, b)
	return k
}
