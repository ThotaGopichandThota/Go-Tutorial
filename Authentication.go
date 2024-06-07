package main

import (
	"errors"
	"fmt"
)

func main() {
	var name, role string
	fmt.Println("Enter the user name:")
	fmt.Scanln(&name)
	fmt.Println("Enter the user role:")
	fmt.Scanln(&role)
	AuthenticatedUser = User{Name: name, Role: role}
	err := AuthenticatedAdminAction()
	if err != nil {
		fmt.Println(err)
	}
}

type User struct {
	Name string
	Role string
}

var AuthenticatedUser User

func AdminOnlyAction() error {
	if AuthenticatedUser.Name == "" {
		return errors.New("not authenticated user")
	}
	if AuthenticatedUser.Role != "admin" {
		return errors.New("permission denied")
	}
	fmt.Println("Admin authenticated successfully")
	return nil
}

func authenticate() error {
	if AuthenticatedUser.Name == "" {
		return errors.New("not authenticated")
	}
	fmt.Println("User", AuthenticatedUser.Name, "authenticated")
	return nil
}

func AuthenticatedAdminAction() error {
	err := authenticate()
	if err != nil {
		return err
	}
	return AdminOnlyAction()
}
