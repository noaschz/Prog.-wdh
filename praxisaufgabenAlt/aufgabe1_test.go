package praxisaufgaben

import "fmt"

func ExampleCommonMultiples() {
	fmt.Println(CommonMultiples(3, 5, 50))
	fmt.Println(CommonMultiples(2, 10, 100))
	fmt.Println(CommonMultiples(1, 1, 10))
	fmt.Println(CommonMultiples(4, 6, 50))
	fmt.Println(CommonMultiples(7, 9, 100))
	fmt.Println(CommonMultiples(5, 5, 25))

}

	// Output:
	// [15 30 45]
	// [10 20 30 40 50 60 70 80 90 100]
	// [1 2 3 4 5 6 7 8 9 10]
	// [12 24 36 48]
	// [63 84]
	// [5 10 15 20 25]

