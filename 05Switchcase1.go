package main

import "fmt"

func main() {
	var month int
	fmt.Println("enter month value")
	fmt.Scanln(&month)
	if month > 0 && month <= 12 {
		switch month {
		case 1:
			fmt.Println("january")
		case 2:
			fmt.Println("february")
		case 3:
			fmt.Println("march")
		case 4:
			fmt.Println("april")
		case 5:
			fmt.Println("may")
		case 6:
			fmt.Println("june")
		case 7:
			fmt.Println("july")
		case 8:
			fmt.Println("august")
		case 9:
			fmt.Println("september")
		case 10:
			fmt.Println("october")
		case 11:
			fmt.Println("noverber")
		case 12:
			fmt.Println("december")
		default:
			fmt.Println("invalid month value")
		}
	} else {
		fmt.Println(" month value should be in between 1-12")
	}
}
