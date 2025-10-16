package main

import "fmt"

func ExampleCountEven() {
	fmt.Println(CountEven([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountEven([]int{7, 9, 11}))
	fmt.Println(CountEven([]int{2, 4, 6, 8}))

	// Output:
	// 3
	// 0
	// 4
}
