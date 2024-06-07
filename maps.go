package main

import "fmt"

func main() {

	b := make(map[int]string)

	b[1] = "Gopichand"
	b[2] = "anil"
	b[3] = "subbu"
	b[4] = "sumanth"
	b[5] = "raj"
	fmt.Println(b)
	//access the element by using key
	fmt.Println(b[4])
	//update the value
	b[2] = "navya"
	fmt.Println(b)
	//add the elements
	b[6] = "anil kumar"
	fmt.Println(b)
	//return the elements one by one
	for k, v := range b {
		fmt.Println(k, v)
	}
	//remove the element from the map
	delete(b, 2)
	fmt.Println(b)
}
