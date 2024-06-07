package main

import (
	"fmt"
	"sort"
)

func main() {
	//as per syntax
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	//create new slice from previous slice
	newSlice := slice[1:4]
	fmt.Println(newSlice)
	//add values to slice
	slice = append(slice, 6, 7, 8, 9, 0)
	fmt.Println(slice)
	//remove the value from the slice
	var coureses = []string{"Gopipchand", "narasimha rao", "bhulakshmi", "kumari", "anusha"}
	fmt.Println(coureses)
	var index int = 2
	coureses = append(coureses[:index], coureses[index+1:]...)
	fmt.Println(coureses)
	//sorting the slice
	sort.Strings(coureses)
	fmt.Println(coureses)
}
