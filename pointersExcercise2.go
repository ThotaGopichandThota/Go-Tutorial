package main

import "fmt"

func main() {

	fmt.Printf("value:%v\t%T\n", a, a)
	fmt.Printf("value:%v\t%T\n", b, b)
	fmt.Printf("value:%v\t%T\n", c, c)
	fmt.Printf("value:%v\t%T\n", d, d)
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

}

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "drop by drop, the bucket is filled"
	q := "persistently, patiently, you are succeed"
	r := "the meaning of life is ........"
	s := 23
	a = &p
	b = &q
	c = &r
	d = &s
}

//in the provided code
