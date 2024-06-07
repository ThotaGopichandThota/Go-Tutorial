package main

import "fmt"

func main() {
	fmt.Println("enter the first swapp number:")
	var a int
	fmt.Scanln(&a)
	fmt.Println("enter the second swapp number:")
	var b int
	fmt.Scanln(&b)
	fmt.Println("before swapping :")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("after swapping:")
	swapp(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
func swapp(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
