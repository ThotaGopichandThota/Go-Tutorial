package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("enter the size of the slice")
	var size int
	fmt.Scanln(&size)
	slice := make([]int, size)
	fmt.Println("enter the slice elements")
	for i := 0; i < size; i++ {
		fmt.Println("the value of index:", i, ":")
		fmt.Scanln(&slice[i])
	}
	fmt.Println("original slice:", slice)
	sort.Ints(slice)
	fmt.Println("the sorted slice:", slice)
	// check the  slice is sorted or not
	fmt.Println(sort.IntsAreSorted(slice))
	// find the length of slice
	fmt.Println("length of slice:", len(slice))

	fmt.Println("_____________________________________")
	navya()
}
func navya() {
	fmt.Println("enter the size of the slice")
	var size int
	fmt.Scanln(&size)
	slice := make([]string, size)
	fmt.Println("enter the slice elements")
	for i := 0; i < size; i++ {
		fmt.Println("enter the value:", i)
		fmt.Scanln(&slice[i])
	}
	fmt.Println("the original slice:", slice)
 
	sort.Strings(slice)
	fmt.Println("sorted slice", slice)
	fmt.Println("the length of slice:", len(slice))
}
