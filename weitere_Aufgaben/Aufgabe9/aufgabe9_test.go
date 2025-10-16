package aufgabe9

import "fmt"

func ExampleMinElements() {
	l1 := []int{0, 2, 0, 2, 0, 2, 0, 2}
	l2 := []int{2, 0, 2, 0, 2, 0, 2, 0}
	l3 := []int{1, 1, 1, 1, 1}

	fmt.Println(MinElements(l1, l2))
	fmt.Println(MinElements(l2, l1))
	fmt.Println(MinElements(l1, l3))

	// Output:
	// [0 0 0 0 0 0 0 0]
	// [0 0 0 0 0 0 0 0]
	// [0 1 0 1 0 2 0 2]
}