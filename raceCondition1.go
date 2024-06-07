package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("wellcome to race condition topic-------------")
	wg := &sync.WaitGroup{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {

		fmt.Println("one 1")
		score = append(score, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		time.Sleep(3 * time.Second)
		fmt.Println("two 2")
		score = append(score, 2)

		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		time.Sleep( 5* time.Second)
		fmt.Println("three 3")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()

	fmt.Println(score)
}
