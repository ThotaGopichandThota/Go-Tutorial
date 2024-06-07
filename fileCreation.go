package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Error :%s", err)
	}
	defer f.Close()
	s := []byte("this is gopi's file")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error:%s", err)
	}
}

// //////READING THE FILE OR  CONTENT IN FILE
func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("the content present in the text file is :", string(dataByte))
}
