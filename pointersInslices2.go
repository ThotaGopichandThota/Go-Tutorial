package main

import "fmt"

func main() {
	fmt.Println("Enter the size of the slice:")
	var size int
	fmt.Scanln(&size)

	fmt.Println("Enter the slice elements:")
	slice := make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter value of index %d: ", i)
		fmt.Scanln(&slice[i])
	}

	fmt.Println("The original slice:", slice)
	modify(slice, size)

	fmt.Println("The modified slice:", slice)
}

func modify(sl []string, size int) {
	var index int
	fmt.Println("Enter the index to modify:")
	fmt.Scanln(&index)
	if index >= 0 && index < size {
		fmt.Printf("Enter the modified value for index %d: ", index)
		fmt.Scanln(&sl[index])
	} else {
		fmt.Println("Invalid index.")
	}
}
