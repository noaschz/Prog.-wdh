package aufgabe10

import "fmt"

func ExampleSumElements() {
	l1 := []int{0, 2, 0, 2, 0, 2, 0, 2}
	l2 := []int{2, 0, 2, 0, 2, 0, 2, 0}
	l3 := []int{1, 1, 1, 1, 1}

	fmt.Println(SumElements(l1, l2))
	fmt.Println(SumElements(l2, l1))
	fmt.Println(SumElements(l1, l3))

	// Output:
	// [2 2 2 2 2 2 2 2]
	// [2 2 2 2 2 2 2 2]
	// [1 3 1 3 1 2 0 2]
}