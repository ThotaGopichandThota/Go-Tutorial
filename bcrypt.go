package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "gopichand@319"
	val:=bcrypt.MinCost
	pass, err := bcrypt.GenerateFromPassword([]byte(password), val)
	if err != nil {
		fmt.Println("error is :", err)
	} else {
		fmt.Println(pass)

	}
	fmt.Println("---------------------------")

}
