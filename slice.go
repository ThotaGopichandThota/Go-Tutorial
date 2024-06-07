package main

import "fmt"

func main() {
	fmt.Println("Enter the size of the slice")
	var size int
	fmt.Scanln(&size)
	slice := make([]int, size)
	fmt.Println("enter the elements")
	for i := 0; i < size; i++ {
		fmt.Println("value of index ", i, ":")
		fmt.Scanln(&slice[i])
	}

	fmt.Println("original slice", slice)
	
}
