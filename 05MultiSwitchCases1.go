package main

import "fmt"

func main() {
	var noOfDays int
	fmt.Println("enter no :")
	fmt.Scanln(&noOfDays)
	if noOfDays >= 28 && noOfDays <= 31 {

		switch noOfDays {
		case 31:
			fmt.Println("january,march,may,july,august,october,december")
		case 30:
			fmt.Println("April,june,september,november")
		case 28, 29:
			fmt.Println("february")
		default:
			fmt.Println("invalid no of days")
		}

	} else {
		fmt.Println("no of days should be in between 28 to 31")
	}
}
