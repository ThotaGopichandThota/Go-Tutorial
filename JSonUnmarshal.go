package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var unmarshal = []byte(`[{"Name":"Gopi","Age":23,"Gender":"male"},{"Name":"navya","Age":22,"Gender":"female"}]`)

	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	var people []Person

	err := json.Unmarshal(unmarshal, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
