package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	wellcome := " wellcome to user input"
	fmt.Println(wellcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give  rating for our biryani:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating :", input)
}
