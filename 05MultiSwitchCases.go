package main

import "fmt"

func main() {
	var day int
	fmt.Println("enter day value:")
	fmt.Scanln(&day)
	if day > 0 && day <= 7 {
		switch day {
		case 1, 3, 5:
			fmt.Println("odd weekday of the week")
		case 2, 4:
			fmt.Println("even weekday of the week")
		case 6, 7:
			fmt.Println("weekend of the week")
		default:
			fmt.Println("invalid day value")
		}
	} else {
		fmt.Println("day value should be in between 1-7")
	}

}
