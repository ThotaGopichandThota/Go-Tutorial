package main

import (
	"fmt"
	"time"
)

func main() {
	go task("gopi")
	go task("chandu")
	time.Sleep(3 * time.Second)
}
func task(name string) {
	fmt.Println(name)
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}
