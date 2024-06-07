package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("wellcome to race condition topic-------------")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("one 1")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("two 2")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("three 3")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
