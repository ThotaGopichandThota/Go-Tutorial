package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := Book{
		Title: "the nightman",
	}
	var n Count = 43
	fmt.Println("the title of the book is :", b)
	fmt.Println("the no.of pages i was read : ", n, "pages")
}

type Book struct {
	Title string
}

func (b Book) String() string {
	title := b.Title
	return title
}

type Count int

func (c Count) string() string {
	count := strconv.Itoa(int(c))
	return count

}
