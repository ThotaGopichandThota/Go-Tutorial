// package main

// import (
// 	"fmt"
// 	"io"
// 	"bytes"
// )

// func main() {
//   p:=People{
// 	Name:"Chandu",
// 	Age: 23,
// 	Gender:"male",

//   }
//   var b bytes.Buffer()
//   p.Gopi(&b)
// }

// type People struct {
// 	Name   string
// 	Age    int
// 	Gender string
// }

//	func (p People) Gopi(w io.Writer) {
//		data := fmt.Sprintf(p.Name, p.Age, p.Gender)
//		w.Write([]byte, data)
//	}
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	p := People{
		Name:   "Chandu",
		Age:    23,
		Gender: "male",
	}
	f, err := os.Create("bagyaRaj")
	if err != nil {
		fmt.Println(err)
	}
	p.Gopi(f)
	var b bytes.Buffer
	p.Gopi(&b)
}

type People struct {
	Name   string
	Age    int
	Gender string
}

func (p People) Gopi(w io.Writer) {
	data := fmt.Sprintf("Name: %s, Age: %d, Gender: %s\n", p.Name, p.Age, p.Gender)
	w.Write([]byte(data))
}
