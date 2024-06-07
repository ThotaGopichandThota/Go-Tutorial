package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var channel = make(chan int)

func main() {
	fmt.Println("main function starts--------------------------")
	wg.Add(2)
	go sendData()
	go receiveData()
	wg.Wait()
	fmt.Println("main function ends----------------------------")
}
func sendData() {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			channel <- i
		}
	}
	close(channel)
	wg.Done()
}
func receiveData() {
	fmt.Println("the odd numbers are :-")
	for {
		if data, ok := <-channel; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	wg.Done()
}
