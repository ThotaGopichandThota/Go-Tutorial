package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var channel = make(chan string)

func main() {
	fmt.Println("main function starts :----------------")
	wg.Add(2)
	go sendData()
	go recieveData()
	wg.Wait()
	fmt.Println("main function ends:--------------------")
}
func sendData() {
	channel <- "my name is Gopichand"
	time.Sleep(1 * time.Second)
	channel <- "iam working in dugong global solutions"
	time.Sleep(2 * time.Second)
	channel <- "iam 23 years old"
	wg.Done()

}
func recieveData() {
	data := <-channel
	fmt.Println("my name is:", data)
	data = <-channel
	fmt.Println("company name:", data)
	data = <-channel
	fmt.Println("age:", data)
	wg.Done()
}
