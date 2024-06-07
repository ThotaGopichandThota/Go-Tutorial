package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter size of an array:")
	fmt.Scanln(&n)
	arr := make([]int, n)
	fmt.Println("Enter the array elements:")
	for i := 0; i < n; i++ {
		//fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Palindrome numbers present in the array:")
	for i := 0; i < len(arr); i++ {
		if isPalindrome(arr[i]) {
			fmt.Println(arr[i])
		}
	}
}

func isPalindrome(num int) bool {
	temp := num
	reverse := 0
	for temp != 0 {
		ld := temp % 10
		reverse = reverse*10 + ld
		temp = temp / 10
	}
	return num == reverse
}
