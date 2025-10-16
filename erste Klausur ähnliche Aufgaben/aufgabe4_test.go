package main

import "fmt"

func ExamplePairwiseSum() {
	fmt.Println(PairwiseSum([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(PairwiseSum([]int{1, 2}, []int{3, 4, 5}))
	fmt.Println(PairwiseSum([]int{}, []int{9, 8, 7}))

	// Output:
	// [5 7 9]
	// [4 6 5]
	// [9 8 7]
}
