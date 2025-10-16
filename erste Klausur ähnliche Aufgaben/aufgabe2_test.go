package main

import "fmt"

func ExampleIncludeBetween() {
	list := []string{"a", "b", "c", "start", "x", "y", "end", "z"}
	fmt.Println(IncludeBetween(list, "start", "end"))
	fmt.Println(IncludeBetween(list, "x", "z"))
	fmt.Println(IncludeBetween(list, "missing", "end"))

	// Output:
	// [x y]
	// [y end]
	// []
}
