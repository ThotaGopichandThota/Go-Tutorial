package main

import "fmt"

func main() {
	var size int
	fmt.Println("enter size of the array")
	fmt.Scanln(&size)
	arr := make([]string, size)
	fmt.Println("enter the array elements")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])

	}
	fmt.Println("the array elements")
	for i := 0; i < size; i++ {
		fmt.Println(arr[i])
	}
}
