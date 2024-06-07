package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter size of an array")
	fmt.Scanln(&n)
	arr := make([]int, n)
	fmt.Println("enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("the palindrome number present in the array: ")
	palindrome(arr)
}
func palindrome(b []int) {
	for i := 0; i < len(b); i++ {

		temp := b[i]
		reverse := 0
		original := temp
		for temp != 0 {
			ld := temp % 10
			reverse = reverse*10 + ld
			temp = temp / 10
		}
		if reverse == original {
			fmt.Println(b[i])
		}
	}
}
