package own

import "fmt"

func ExampleFilterBetween_standard_cases() {
	list := []int{1, 5, 10, 15, 20, 25, 30}
	m1, n1 := 5, 20
	m2, n2 := 10, 30
	m3, n3 := 0, 100

	fmt.Println(FilterBetween(list, m1, n1))
	fmt.Println(FilterBetween(list, m2, n2))
	fmt.Println(FilterBetween(list, m3, n3))

	// Output:
	// [10 15]
	// [15 20 25]
	// [1 5 10 15 20 25 30]
}

func ExampleFilterBetween_special_cases() {
	list := []int{3, 6, 9, 12, 15}
	m1, n1 := 7, 11  // Kein Wert liegt zwischen 7 und 11
	m2, n2 := 3, 15  // Alle Werte außer 3 und 15 sollten enthalten sein
	m3, n3 := 16, 20 // Keine Werte im Bereich

	fmt.Println(FilterBetween(list, m1, n1)) // []
	fmt.Println(FilterBetween(list, m2, n2)) // [6 9 12]
	fmt.Println(FilterBetween(list, m3, n3)) // []

	// Output:
	// [9]
	// [6 9 12]
	// []
}
