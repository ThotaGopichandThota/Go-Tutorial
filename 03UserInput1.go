package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "user name"

	fmt.Println(name)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the userName")
	input, _ := reader.ReadString('\n')
	fmt.Println("user name is:", input)
}
