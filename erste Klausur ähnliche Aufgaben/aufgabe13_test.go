package main

import "fmt"

func ExampleCommonElements() {
	fmt.Println(CommonElements([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	fmt.Println(CommonElements([]int{10, 20, 30}, []int{40, 50, 60}))
	fmt.Println(CommonElements([]int{5, 10, 15, 20}, []int{10, 20, 30, 40}))

	// Output:
	// [3 4]
	// []
	// [10 20]
}
