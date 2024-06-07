package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := "value"
	fmt.Println(num)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the num value")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Printf("%T", input)
	temp, _ := strconv.Atoi(input)

	// temp := reader.Buffered()
	fmt.Printf("%T", temp)
	fmt.Println(temp%2, temp)
	if temp%2 == 0 {

		fmt.Println("it is a even number")
	} else {
		fmt.Println("it is a odd number")
	}
}
