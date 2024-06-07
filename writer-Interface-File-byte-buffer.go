package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	p := Person{
		Name:   "Gopichand",
		Age:    23,
		Gender: "male",
	}
	f, err := os.Create("Gopi1111.txt")
	if err != nil {
		log.Fatalf("Error:%s", err)
	}
	defer f.Close()

	var b bytes.Buffer
	p.Gopi(f)
	p.Gopi(&b)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) Gopi(w io.Writer) {
	data := fmt.Sprintf("Name: %s, Age:%d,Gender:%s\n", p.Name, p.Age, p.Gender)
	w.Write([]byte(data))
}
