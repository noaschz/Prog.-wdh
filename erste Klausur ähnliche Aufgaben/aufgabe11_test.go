package main

import "fmt"

func ExampleRemoveMultiples() {
	fmt.Println(RemoveMultiples([]int{3, 6, 9, 12, 15, 18}, 3))
	fmt.Println(RemoveMultiples([]int{10, 20, 30, 40, 50}, 5))
	fmt.Println(RemoveMultiples([]int{1, 2, 3, 4, 5, 6}, 2))

	// Output:
	// []
	// []
	// [1 3 5]
}
