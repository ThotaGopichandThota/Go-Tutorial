package main

import "fmt"

func main() {
	fmt.Println("enter the size of the slice:")
	var size int
	fmt.Scanln(&size)
	slice := make([]int, size)
	fmt.Println("enter the slice elements:")
	for i := 0; i < size; i++ {
		//fmt.Println(i)
		fmt.Scanln(&slice[i])
	}
	fmt.Println("the normal order of slice is:", slice)
	fmt.Println("the reverse oeder of the slice :")
	for i := size - 1; i >= 0; i-- {
		//fmt.Println("the reverse order of  the slice is:", slice[i])

		fmt.Println(slice[i])
	}

}
