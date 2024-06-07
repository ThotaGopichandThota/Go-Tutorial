package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex

func main() {
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5}
	result := []int{}
	fmt.Println("original slice is :", input)
	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}
	wg.Wait()
	fmt.Println("resultant slice is:", result)
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	mut.Lock()
	defer wg.Done()
	processData := data * 2
	*result = append(*result, processData)
	mut.Unlock()
}
