package main

import "fmt"

func main() {
	fmt.Println("wellcome to structs topic")
	// user one
	Gopichand := User{"Gopichand ", "tgopichand74@gmail.com", true, 24}
	fmt.Println(Gopichand)

	fmt.Printf("Gopichand details are: %+v \n", Gopichand)
	// print the particular parameters
	fmt.Printf("name is :%v  and mail is:%v", Gopichand.Name, Gopichand.Email)

	//user two
	anil := User{"anil", "anil123@gmail.com", true, 25}
	fmt.Print(anil)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
