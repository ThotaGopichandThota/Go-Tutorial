package main

import "fmt"

func main() {
	func() {
		fmt.Println("wellcome to anonymous functions")
	}()
	gopi()
	fmt.Println("execution end of anonymous function")
}
func gopi() {
	func() {
		var num int
		fmt.Println("enter the num value: ")
		fmt.Scanln(&num)
		reverse := 0
		temp := num
		for temp != 0 {
			ld := temp % 10
			reverse = reverse*10 + ld
			temp = temp / 10
		}
		if reverse == num {
			fmt.Println("it is a palinfrome number")

		} else {
			fmt.Println("it is not a palindrome number")
		}
	}()
}
