package own

import "fmt"

func ExampleCountX_standard_cases() {
	list := []int{1, 2, 3, 4, 2, 2, 5, 2}
	x1 := 2
	x2 := 3
	x3 := 6

	fmt.Println(CountX(list, x1))
	fmt.Println(CountX(list, x2))
	fmt.Println(CountX(list, x3))

	// Output:
	// 4
	// 1
	// 0
}

func ExampleCountX_special_cases() {
	list := []int{7, 7, 7, 7, 7}
	x1 := 7
	x2 := 1
	emptyList := []int{}

	fmt.Println(CountX(list, x1))     // Alle Elemente sind 7
	fmt.Println(CountX(list, x2))     // 1 kommt nicht vor
	fmt.Println(CountX(emptyList, 5)) // Leere Liste → 0

	// Output:
	// 5
	// 0
	// 0
}
