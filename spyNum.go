package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("enter the start limit: ")
	fmt.Scanln(&start)
	fmt.Println("enter the end limit")
	fmt.Scanln(&end)
	// sum:=0
	// product :=1
	for i := start; i < end; i++ {
		sum := 0
		product := 1
		temp := i
		for temp != 0 {
			ld := temp % 10
			sum = sum + ld
			product = product * ld
			temp = temp / 10
		}
		if sum == product {
			fmt.Println(i)
		}
	}

}
