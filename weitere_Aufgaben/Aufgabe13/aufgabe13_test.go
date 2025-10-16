package aufgabe13

import "fmt"

func ExampleUnion() {
	l1 := []int{20, 0, 15, 5, 10}
	l2 := []int{0, 10, 20, 30, 40}
	l3 := []int{15, 25, 5, -5, 35, 5, 25}

	fmt.Println(Union(l1, l2))
	fmt.Println(Union(l1, l3))
	fmt.Println(Union(l2, l3))

	// Output:
	// [20 0 15 5 10 30 40]
	// [20 0 15 5 10 25 -5 35]
	// [0 10 20 30 40 15 25 5 -5 35]
}
