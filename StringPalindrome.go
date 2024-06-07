package main

import (
	"fmt"
)

func main() {
	var original string
	fmt.Println("enter the string :")
	fmt.Scanln(&original)
	reverse := ""
	for i := len(original) - 1; i >= 0; i-- {
		ch := []rune(original)
		reverse = reverse + string(ch[i])
	}
	if reverse == original {
		fmt.Println("it is a palindrome string")
	} else {
		fmt.Println("it is not palindrome string")
	}

}
