package own

import "fmt"

func ExampleDuplicates() {
	list1 := []int{1, 2, 3, 4, 2, 3, 5, 6, 3}
	list2 := []int{7, 8, 9, 10}
	list3 := []int{1, 1, 1, 2, 2, 3, 3, 3}
	list4 := []int{}

	fmt.Println(Duplicates(list1))
	fmt.Println(Duplicates(list2))
	fmt.Println(Duplicates(list3))
	fmt.Println(Duplicates(list4))

	// Output:
	// [2 3]
	// []
	// [1 2 3]
	// []
}
