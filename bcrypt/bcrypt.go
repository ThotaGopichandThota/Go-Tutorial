package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "gopichand@319"
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error is :", err)
	}
	fmt.Println(pass)

	password1 := `gopichand@319`
	err = bcrypt.CompareHashAndPassword(pass, []byte([]byte(password1)))

	if err != nil {
		fmt.Println("login failed")
		return
	}

	fmt.Println("login succeessful")
}
