/*

func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/

package main

import "fmt"

func main() {
	// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

	fmt.Println("wellcome to functions topic")
	//function calling
	myName()
	info("gopichand", 23)
	info("NarasimhaRao", 52)
	info("bhulakshmi", 45)
	fmt.Println("the sum of the integer :", addition(3, 6))
	fmt.Println("the multiplication of number is :", multiplication(1, 2, 3, 4))
	s1,n:=details("chandu",23)
	fmt.Println(s1,n)
}

// funtion with no argument
func myName() {
	fmt.Println("my name is Gopichand")
}

// function with no.of arguments
func info(name string, age int) {
	fmt.Println(" name is:", name, age, "years old")

}

// return the  arguments
func addition(a int, b int) int {
	total := a + b
	return total

}

// variadic function
func multiplication(num ...int) int {
	total := 1
	for _, num := range num {
		total = total * num
	}
	return total
}
func details(name1 string,age int)string,int{
	age *=2
	return name1,age
}
