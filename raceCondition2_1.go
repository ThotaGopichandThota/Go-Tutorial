package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))
	fmt.Println("originalslice", input)
	for i, data := range input {

		wg.Add(1)
		processData(&wg, &result[i], data)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println("resulted slice is:", result)
}
func processData(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done()

	processedData := process(data)
	*resultDest = processedData
}
func process(data int) int {
	time.Sleep(2 * time.Second)
	return data * 2
}
