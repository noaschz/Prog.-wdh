package main

import "fmt"

func ExampleMaxInList() {
	fmt.Println(MaxInList([]int{1, 3, 7, 2, 5}))
	fmt.Println(MaxInList([]int{10, 20, 30, 40}))
	fmt.Println(MaxInList([]int{-5, -10, -1}))

	// Output:
	// 7
	// 40
	// -1
}
