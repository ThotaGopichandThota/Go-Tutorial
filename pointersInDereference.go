package main

import "fmt"

func inDirect(n *int) {
	*n = 55
}
func main() {
	a := 22
	fmt.Println(a)
	inDirect(&a)
	fmt.Println(a)
	//slice
	xi := []int{1, 3, 4, 5, 3, 0}
	//before calling dereference pointer
	fmt.Println(xi)
	//calling dereference pointer
	slicedelta(&xi)
	fmt.Println("after calling dereferece pointer:", xi)
	// map

	mp := make(map[int]string)
	mp[1] = "gopichand"
	fmt.Println(mp[1])
	mapDelta(&mp, 1)
	fmt.Println("after calling dereferece pointer:", mp[1])

}

// slice

func slicedelta(ni *[]int) {
	(*ni)[0] = 23
	(*ni)[1] = 33
}

// map
func mapDelta(md *map[int]string, n int) {
	(*md)[n] = "navy"
}
