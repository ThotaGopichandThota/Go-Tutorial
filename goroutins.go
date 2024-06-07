package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	 go gopi()
	chandu()
	wg.Wait()
}
func gopi() {
	for i := 0; i < 9; i++ {
		fmt.Println("hiii Goood afternoon")
	}
	wg.Done()
}
func chandu() {
	for i := 0; i < 6; i++ {
		fmt.Println("good morining............")
	}
}
