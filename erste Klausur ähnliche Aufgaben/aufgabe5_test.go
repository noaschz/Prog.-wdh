package main

import "fmt"

func ExampleUniqueElements() {
	fmt.Println(UniqueElements([]int{1, 2, 2, 3, 4, 4, 5}))
	fmt.Println(UniqueElements([]int{10, 20, 20, 30, 40, 10}))
	fmt.Println(UniqueElements([]int{7, 8, 9, 7, 8, 9}))

	// Output:
	// [1 3 5]
	// [30 40]
	// []
}
