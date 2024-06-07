package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := readFile("nana1.txt")
	if err != nil {
		log.Fatalf("read file in main:%s", err)
	}
	fmt.Println(f)
	fmt.Println(string(f))
}

func readFile(fileName string) ([]byte, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("errors in readFile is:%s", err)
	}
	return f, nil
}
