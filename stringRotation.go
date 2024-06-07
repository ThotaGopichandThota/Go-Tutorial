package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string:")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	fmt.Println("Enter the number of rotations:")
	var rotations int
	fmt.Scanln(&rotations)

	fmt.Println("Original string:", str)
	fmt.Println("Rotated string:", RotateString(str, rotations))
}

// RotateString rotates a string s by n positions to the left
func RotateString(s string, n int) string {
	length := len(s)
	if length == 0 {
		return s
	}
	n = n % length // Normalize n to be within the range of string length
	return s[n:] + s[:n]
}
