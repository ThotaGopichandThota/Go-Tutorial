package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time handling concept")
	//present time running on your laptop
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2024-4-9, wednesday"))
	//create our own time
	createTime := time.Date(1999, time.October, 9, 4, 0, 0, 0, time.Local)
	fmt.Println(createTime)
	fmt.Println(createTime.Format("2024-04-9,monday"))
}
