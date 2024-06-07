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
	pointer := &slice
	sliceDelta(*pointer)
	fmt.Println("modified slice", slice)
}

func sliceDelta(sl []int) {
	// sl[1] = 4
	// sl[3] = 8
	fmt.Println("enter the value of index2:")
	fmt.Scanln(&sl[2])
	fmt.Println("enter the value of index3:")
	fmt.Scanln(&sl[3])
}
