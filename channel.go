package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var channel = make(chan string)

func main() {
	fmt.Println("main function starts :")
	wg.Add(2)
	go sendData()
	go recieveData()
	wg.Wait()

	fmt.Println("main function ends")

}
func sendData() {
	channel <- "hii good afternoon"
	wg.Done()
}
func recieveData() {
	data := <-channel
	fmt.Println("receive the data from the sendData:", data)
	wg.Done()
}
