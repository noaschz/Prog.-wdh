package aufgabe14

import "fmt"

func ExampleIntersection() {
	l1 := []int{20, 0, 15, 5, 10}
	l2 := []int{0, 10, 20, 30, 40}
	l3 := []int{15, 25, 5, -5, 35, 5, 25}

	fmt.Println(Intersection(l1, l2))
	fmt.Println(Intersection(l1, l3))
	fmt.Println(Intersection(l2, l3))

	// Output:
	// [20 0 10]
	// [15 5]
	// []
}

