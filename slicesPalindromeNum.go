package main

import "fmt"

func main() {
	fmt.Println("Enter the size of slice:")
	var size int
	fmt.Scanln(&size)
	slice := make([]int, size)
	fmt.Println("Enter the slice elements:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&slice[i])
	}
	fmt.Println("Original slice:", slice)
	fmt.Println("the palindrome numbers present in the slice:")
	palindrome(slice)
	fmt.Println("the spy numbers present in the slice :")
	spy(slice)
}

func palindrome(n []int) {
	for _, num := range n {
		reverse := 0
		temp := num

		for temp != 0 {
			ld := temp % 10
			reverse = reverse*10 + ld
			temp /= 10
		}
		if reverse == num {
			fmt.Println(num)
		}
	}
}

func spy(n []int) {
	for _, num := range n {
		sum := 0
		product := 1
		temp := num

		for temp != 0 {
			ld := temp % 10
			sum = sum + ld
			product = product * ld
			temp = temp / 10
		}
		if sum == product {
			fmt.Println(num)
		}
	}
}
