package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var channel = make(chan int)

func main() {
	fmt.Println("main function starts:-----------")
	wg.Add(2)
	go sendData()
	go receiveData()

	wg.Wait()
	fmt.Println("main function ends:-------------")
}

func sendData() {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			channel <- i
		}
	}
	close(channel)
}

func receiveData() {
	defer wg.Done()
	for data := range channel {
		fmt.Println(data)
	}
}
