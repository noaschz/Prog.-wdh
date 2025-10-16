package aufgabe4

import "fmt"

func ExampleRecursiveMaxElements() {
	l1 := []int{0, 2, 0, 2, 0, 2, 0, 2}
	l2 := []int{2, 0, 2, 0, 2, 0, 2, 0}
	l3 := []int{1, 1, 1, 1, 1}

	fmt.Println(RecursiveMaxElements(l1, l2))
	fmt.Println(RecursiveMaxElements(l2, l1))
	fmt.Println(RecursiveMaxElements(l1, l3))

	// Output:
	// [2 2 2 2 2 2 2 2]
	// [2 2 2 2 2 2 2 2]
	// [1 2 1 2 1 2 0 2]
}