package main

import "fmt"

func main() {
	gopi := 42
	fmt.Printf("42 num convert to binary,%b \n", gopi)
	fmt.Printf("42 num conver to hexadecimal,%x \n", gopi)

	a, b, c, d := 1, 2, 3, 4
	fmt.Print("%v \t %b \t %x \n", a, a, a)
	fmt.Printf("%v \t %b \t %x \n", b, b, b)
	fmt.Printf("%v \t %b \t %x \n", c, c, c)
	fmt.Printf("%v \t %b \t %x \n", d, d, d)
}
