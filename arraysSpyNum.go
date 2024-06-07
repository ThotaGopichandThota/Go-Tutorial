package main

import "fmt"

func main() {
	var size int
	fmt.Println("enter the size of array: ")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("enter the array elements:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < len(arr); i++ {
		if spy(arr[i]) {
			fmt.Println("the spy numbers present in the array:", arr[i])
		}
	}
}
func spy(num int) bool {
	temp := num
	sum := 0
	product := 1
	for temp != 0 {
		ld := temp % 10
		sum = sum + ld
		product = product * ld
		temp = temp / 10

	}
	return sum == product

}
