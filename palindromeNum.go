package main

import "fmt"

func main() {

	var start, end int
	fmt.Println("enter the starting limit")
	fmt.Scanln(&start)
	fmt.Println("enter the end limit")
	fmt.Scanln(&end)
	count := 0
	for i := start; i <= end; i++ {
		reverse := 0
		temp := i
		for temp != 0 {
			ld := temp % 10
			reverse = reverse*10 + ld
			temp /= 10
		}
		if reverse == i {
			fmt.Println(i)
			count++
		}

	}
	fmt.Println("the count of palindrome numbers in between the range: ", count)
}
