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
		Name:   "NarasimhaRao",
		Age:    52,
		Gender: "male",
	}
	f, err := os.Create("nana1.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()

	p.Gopi(f)

	var b bytes.Buffer
	p.Gopi(&b)

	fmt.Println(b.String())
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) Gopi(w io.Writer) {
	data := fmt.Sprintf("Name:%s, Age:%d, Gender:%s\n", p.Name, p.Age, p.Gender)
	w.Write([]byte(data))
}
