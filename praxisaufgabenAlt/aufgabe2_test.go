package praxisaufgaben

import "fmt"

func ExampleArraySums() {
	fmt.Println(ArraySums([]int{1, 3, 5, 7}))
	fmt.Println(ArraySums([]int{1, 1, 2, 80}))
	fmt.Println(ArraySums([]int{7, 3, 1, 2}))
	fmt.Println(ArraySums([]int{3, 3, 0, 2}))
	fmt.Println(ArraySums([]int{1, 2, 3, 4, 5}))

	// Output:
	// [1 4 9 16]
	// [1 2 4 84]
	// [7 10 11 13]
	// [3 6 6 8]
	// [1 3 6 10 15]
}
