package main

import "fmt"

func main() {

	var num int
	fmt.Println("enter number:")
	fmt.Scanln(&num)

	palindrome(num)
	primeNum(num)
}

func palindrome(original int) {
	revesre := 0
	temp := original
	for temp != 0 {
		ld := temp % 10
		revesre = revesre*10 + ld
		temp = temp / 10
	}
	if revesre == original {
		fmt.Println("it is palindrome number")
	} else {
		fmt.Println("it is not a palindrome number")
	}

}
func primeNum(prime int) {
	temp := prime
	status := true
	for i := 2; i < temp; i++ {
		if temp%i == 0 {
			status = false
			break
		}
	}
	if status {

		fmt.Println("it is a prime number")

	} else {
		fmt.Println("it is not a prime number")
	}
}
