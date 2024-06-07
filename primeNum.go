package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("enter the start limit value: ")
	fmt.Scanln(&start)
	fmt.Println("enter the end limit value: ")
	fmt.Scanln(&end)

	for i := start; i <= end; i++ {
		status := true
		temp := i
		for j := 2; j < temp; j++ {
			if temp%j == 0 {
				status = false
				break
			}
		}
		if status {
			if i > 1 {
				fmt.Println(i)
			}

		}
	}

}
