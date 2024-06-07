package main

import "fmt"

func main() {
	var day int
	fmt.Println("enter day value :")
	fmt.Scanln(&day)
	if day > 0 && day <= 7 {
		switch day {
		case 1:
			fmt.Println("sunday")
		case 2:
			fmt.Println("monday")
		case 3:
			fmt.Println("tuesday")
		case 4:
			fmt.Println("wednesday")
		case 5:
			fmt.Println("thursday")
		case 6:
			fmt.Println("friday")
		case 7:
			fmt.Println("saturday")
		default:
			fmt.Println("invalid day value")
		}

	} else {
		fmt.Println("the day value in between 1-7")
	}
}
