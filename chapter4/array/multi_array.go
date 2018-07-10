package main

import "fmt"

func main() {
	var marray1 [4][2]int
	marray2 := [4][2]int {{10, 11}, {20, 21},{30, 31},{40, 41}}
	marray3 := [4][2]int {1: {20, 21}, 3: {40, 41}}
	marray4 := [4][2]int {1: {0: 20}, 3: {1:41}}

	fmt.Println(marray1)
	fmt.Println(marray2)
	fmt.Println(marray3)
	fmt.Println(marray4)

	var marray5 [1e6]int

	foo(&marray5)
}

func foo(array *[1e6]int) {

}
