package main

import "fmt"

func ExampleMergeSortedListsRecursive_standard_cases() {
	l1 := []int{1, 3, 5}
	l2 := []int{2, 4, 6}
	l3 := []int{7, 8, 9}
	l4 := []int{}

	fmt.Println(MergeSortedListsRecursive(l1, l2))
	fmt.Println(MergeSortedListsRecursive(l1, l3))
	fmt.Println(MergeSortedListsRecursive(l1, l4))

	// Output:
	// [1 2 3 4 5 6]
	// [1 3 5 7 8 9]
	// [1 3 5]
}
