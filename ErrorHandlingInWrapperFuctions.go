package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("enter the first value")
	var a float64
	fmt.Scanln(&a)
	fmt.Println("enter the second value")
	var b float64
	fmt.Scanln(&b)
	result, err := safety(a, b)
	if err != nil {
		fmt.Println("safe divide :%w", err)
	} else {
		fmt.Println("result :%d", result)
	}
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("the given number is divisible by zero:")
	}
	return a / b, nil
}
func safety(a, b float64) (float64, error) {
	result, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("safe divident", err)
	}
	return result, nil
}
