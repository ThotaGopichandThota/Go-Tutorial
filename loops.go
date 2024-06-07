package main

import "fmt"

func main() {
	count := 0
	product := 1
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i)
			count++
			product = product * i
		}

	}
	fmt.Println("the no of even numbers present in the range ", count)
	fmt.Println("the product of even numbers in the range is: ", product)
}
