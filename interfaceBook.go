package main

import "fmt"

func main() {
	var pages int
	fmt.Println("enter the no.of pages ")
	fmt.Scanln(&pages)
	book := MyBook{pages: pages}
	Write(book)
	Read(book)
}

type Book interface {
	Read() int
	Write() int
}
type MyBook struct {
	pages int
}

func (b MyBook) Read() int {
	reading := b.pages
	return reading
}
func (b MyBook) Write() int {
	writting := b.pages
	return writting
}
func Read(b Book) {
	fmt.Println("reading :", b.Read(), "pages per day")
}
func Write(b Book) {
	fmt.Println("writting :", b.Write(), "pages per day")
}
