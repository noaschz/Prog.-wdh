package own

import "fmt"

func ExampleProductOfLists_standard_cases() {
	l1 := []int{2, 3, 4, 5}
	l2 := []int{1, 2, 3, 4}

	fmt.Println(ProductOfLists(l1, l2))

	l3 := []int{1, 2, 3}
	l4 := []int{4, 5, 6}

	fmt.Println(ProductOfLists(l3, l4))

	// Output:
	// [2 6 12 20]
	// [4 10 18]
}

func ExampleProductOfLists_empty_cases() {
	l1 := []int{2, 3, 4, 5}
	l2 := []int{}

	fmt.Println(ProductOfLists(l1, l2))
	fmt.Println(ProductOfLists(l2, l1))
	fmt.Println(ProductOfLists(l2, l2))

	// Output:
	// [2 3 4 5]
	// [2 3 4 5]
	// []
}

func ExampleProductOfLists_different_length_cases() {
	l1 := []int{2, 3, 4}
	l2 := []int{1, 2}

	fmt.Println(ProductOfLists(l1, l2))

	l3 := []int{1, 2}
	l4 := []int{2, 3, 4}

	fmt.Println(ProductOfLists(l3, l4))

	l5 := []int{5, 6, 7, 8}
	l6 := []int{2, 3}

	fmt.Println(ProductOfLists(l5, l6))

	// Output:
	// [2 6 4]
	// [2 6 4]
	// [10 18 7 8]
}
