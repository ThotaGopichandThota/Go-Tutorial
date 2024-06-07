package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal = []string{}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websitelist := []string{
		"https://googlge.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
}
func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println(signal, endPoint)
	} else {
		mut.Lock()
		signal = append(signal, endPoint)
		mut.Unlock()
		fmt.Printf("%d statuscode for %s\n", res.StatusCode, endPoint)
	}
}
