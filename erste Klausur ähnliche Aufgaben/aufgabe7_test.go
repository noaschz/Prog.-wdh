package main

import "fmt"

func ExampleMinInList() {
	fmt.Println(MinInList([]int{1, 3, 7, 2, 5}))
	fmt.Println(MinInList([]int{10, 20, 30, 40}))
	fmt.Println(MinInList([]int{-5, -10, -1}))

	// Output:
	// 1
	// 10
	// -10
}
