package main

import "fmt"

func main() {
	// anonymous function is assigning to a variable is fuction expression
	temp := func() {
		fmt.Println("iam gopichand")
	}

	temp()
	//anonymous function with arguments assigning to variable
	temp1 := func(s string) {
		fmt.Println("iam from :", s)
	}

	temp1("Andhra pradesh")

}

func (){
	fmt.Println("my father name is NarasimhaRao")
}()
