package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(jointer("my", "name", "is", "gopichand", "-"))
	fmt.Println("the count numbers in range:", count(1, 2, 5, 7))
}
func jointer(name ...string) string {
	return strings.Join(name, "_")
}
func count(num ...int) int {

	return len(num)
}
