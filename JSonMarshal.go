package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	p := Person{
		Name:   "Gopi",
		Age:    23,
		Gender: "male",
	}
	p2 := Person{
		Name:   "navya",
		Age:    22,
		Gender: "female",
	}

	People := []Person{p, p2}
	fmt.Println(People)
	//creating a marshal file
	marshal, err := json.Marshal(People)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(marshal) //bytecode
	fmt.Println(string(marshal))

}
