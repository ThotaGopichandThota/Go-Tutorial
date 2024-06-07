package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)
}
