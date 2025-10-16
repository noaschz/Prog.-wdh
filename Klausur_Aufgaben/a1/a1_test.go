package a1

import "fmt"

func ExampleIncludeStringsBetween() {
	l1 := []string{"abc", "BEGIN", "def", "END", "ghi"}
	l2 := []string{"BEGIN", "def", "END", "ghi"}
	l3 := []string{"BEGIN", "def", "ghi", "END"}
	l4 := []string{"END", "def", "BEGIN", "ghi"}
	l5 := []string{"abc", "def", "ghi"}

	fmt.Println(IncludeStringsBetween(l1, "BEGIN", "END"))
	fmt.Println(IncludeStringsBetween(l2, "BEGIN", "END"))
	fmt.Println(IncludeStringsBetween(l3, "BEGIN", "END"))
	fmt.Println(IncludeStringsBetween(l4, "BEGIN", "END"))
	fmt.Println(IncludeStringsBetween(l5, "BEGIN", "END"))

	// Output:
	// [def]
	// [def]
	// [def ghi]
	// []
	// []
}