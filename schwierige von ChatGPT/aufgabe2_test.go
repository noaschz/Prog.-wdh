package main

import "fmt"

func ExampleSumOfEvensRecursive_standard_cases() {
	list1 := []int{1, 2, 3, 4, 5, 6}
	list2 := []int{7, 9, 11}

	fmt.Println(SumOfEvensRecursive(list1))
	fmt.Println(SumOfEvensRecursive(list2))

	// Output:
	// 12
	// 0
}
