package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "http://loc.dev"

func main() {
	fmt.Println("wellcome to web request topic:")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response:%T\n", response)
	defer response.Body.Close() //developer responsibility for calling the this line

	dataBytes, err := ioutil.ReadAll(response.Body)
	//dataBytes,err:=ioutil.ReadFile(.txt)
	if err != nil {
		fmt.Println("the error is:", err)
	}
	fmt.Println(dataBytes)

	content := string(dataBytes)
	fmt.Println(content)
}
