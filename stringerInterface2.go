package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	book := Book{
		Title: "the night man",
		Price: 340,
	}
	var n Count = 34
	//log.Println(book)
	loginfo(book)
	log.Println("the no of  pages i was read: ", n, "pages")
}

type Book struct {
	Title string
	Price float64
}
type Count int

func (c Count) String() string {
	count := strconv.Itoa(int(c))
	return count
}

func (b Book) String() string {
	return fmt.Sprintf("Title:%s  Price:%.2f", b.Title, b.Price)
}

func loginfo(s fmt.Stringer) {
	log.Println("iam loging from dugong startup company", s.String())
}
