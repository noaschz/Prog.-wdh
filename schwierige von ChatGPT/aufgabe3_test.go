package main

import "fmt"

func ExampleFindUniqueElements_standard_cases() {
	list1 := []int{1, 2, 2, 3, 4, 4, 5}
	list2 := []int{7, 7, 7, 8, 9, 10, 10}

	fmt.Println(FindUniqueElements(list1))
	fmt.Println(FindUniqueElements(list2))

	// Output:
	// [1 3 5]
	// [8 9]
}
