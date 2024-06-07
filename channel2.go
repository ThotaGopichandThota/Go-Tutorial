package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var channel = make(chan string, 6) // 6 value is the buffer memory length

func main() {
	fmt.Println("the main function starts:-----------------------")
	wg.Add(2)
	go sendData()
	go receiveData()
	wg.Wait()

	fmt.Println("the main function ends:-------------------------")
}

func sendData() {
	defer wg.Done()
	channel <- "B.Tech"
	time.Sleep(1 * time.Second)
	channel <- "Mechanical"
	time.Sleep(2 * time.Second)
	channel <- "Guntur Engineering College"

	wg.Done()
}

func receiveData() {
	data := <-channel
	fmt.Println("qualifaication:-", data)
}
